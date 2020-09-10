import NoSsr from '@material-ui/core/NoSsr';
import React from 'react';
import { Controlled as CodeMirror } from 'react-codemirror2';
import {
  withStyles, Grid, TextField, IconButton, Dialog, DialogTitle, DialogContent, DialogActions, Divider, Card, CardHeader, CardActions, Menu, MenuItem, Chip
} from '@material-ui/core';
import { blue } from '@material-ui/core/colors';
import PropTypes from 'prop-types';
import { withRouter } from 'next/router';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import CloseIcon from '@material-ui/icons/Close';
import { withSnackbar } from 'notistack';
import AddIcon from '@material-ui/icons/Add';
import DeleteIcon from '@material-ui/icons/Delete';
import PlayIcon from '@material-ui/icons/PlayArrow';
import { updateProgress } from '../lib/store';
import dataFetch from '../lib/data-fetch';
import MUIDataTable from "mui-datatables";

const styles = (theme) => ({
  root: {
    padding: theme.spacing(10),
    width: '100%',
  },
  chipGrid: {
    padding: theme.spacing(10),
    width: '100%',
    paddingBottom: '0',
  },
  buttons: {
    width: '100%',
  },
  button: {
    marginTop: theme.spacing(3),
    marginLeft: theme.spacing(1),
  },
  margin: {
    margin: theme.spacing(1),
  },
  alreadyConfigured: {
    textAlign: 'center',
    padding: theme.spacing(20),
  },
  chip: {
    height: '40px',
    marginRight: theme.spacing(5),
    marginTop: theme.spacing(4),
    marginBottom: theme.spacing(-5),
    fontSize: '15px',
  },
  colorSwitchBase: {
    color: blue[300],
    '&$colorChecked': {
      color: blue[500],
      '& + $colorBar': {
        backgroundColor: blue[500],
      },
    },
  },
  colorBar: {},
  colorChecked: {},
  uploadButton: {
    margin: theme.spacing(1),
    marginTop: theme.spacing(3),
  },
  fileLabel: {
    width: '100%',
  },
  editorContainer: {
    width: '100%',
  },
  deleteLabel: {
    paddingRight: theme.spacing(2),
  },
  alignRight: {
    textAlign: 'right',
  },
  alignLeft: {
    textAlign: 'left',
    marginLeft: theme.spacing(1),
  },
  padLeft: {
    paddingLeft: theme.spacing(0.25),
  },
  padRight: {
    paddingRight: theme.spacing(0.25),
  },
  deleteRight: {
    float: 'right',
  },
  expTitleIcon: {
    width: theme.spacing(3),
    display: 'inline',
    verticalAlign: 'middle',
  },
  expIstioTitleIcon: {
    width: theme.spacing(2),
    display: 'inline',
    verticalAlign: 'middle',
    marginLeft: theme.spacing(0.5),
    marginRight: theme.spacing(0.5),
  },
  expTitle: {
    display: 'inline',
    verticalAlign: 'middle',
  },
  icon: {
    width: theme.spacing(2.5),
  },
});


class MesheryAdapterPlayComponent extends React.Component {

  
  constructor(props) {
    super(props);

    this.cmEditorAdd = null;
    this.cmEditorDel = null;

    const { adapter } = props;

    const menuState = {};

    this.addIconEles = {};
    this.delIconEles = {};
    // initializing menuState;
    if (adapter && adapter.ops) {
      // NOTE: this will have to updated to match the categories
      [0, 1, 2, 3, 4].forEach((i) => {
        menuState[i] = {
          add: false,
          delete: false,
        };
      });
    }

    this.state = {
      selectedOp: '',
      cmEditorValAdd: '',
      cmEditorValAddError: false,

      cmEditorValDel: '',
      cmEditorValDelError: false,

      selectionError: false,

      namespace: 'default',
      namespaceError: false,

      customDialogAdd: false,
      customDialogDel: false,
      customDialogSMI: false,

      open : false,

      menuState, // category: {add: 1, delete: 0}
    };
  }

  handleChange = (name, isDelete = false) => {
    const self = this;
    return (event) => {
      if (name === 'namespace' && event.target.value !== '') {
        this.setState({ namespaceError: false });
      }

      if (name === 'selectedOp' && event.target.value !== '') {
        if (event.target.value === 'custom') {
          if (isDelete) {
            if (self.state.cmEditorValDel !== '' && self.cmEditorDel.state.lint.marked.length === 0) {
              self.setState({ selectionError: false, cmEditorValDelError: false });
            }
          } else if (self.state.cmEditorValAdd !== '' && self.cmEditorAdd.state.lint.marked.length === 0) {
            self.setState({ selectionError: false, cmEditorValAddError: false });
          }
        } else {
          self.setState({ selectionError: false });
        }
      }

      self.setState({ [name]: event.target.value });
    };
  }

  handleModalClose(isDelete) {
    const self = this;
    return () => {
      const item = isDelete ? 'customDialogDel' : 'customDialogAdd';
      self.setState({ [item]: false });
    };
  }

  handleSMIClose() {
    const self = this;
    return () => {
      self.setState({['customDialogSMI']: false });
    }
  }

  handleModalOpen(isDelete) {
    const self = this;
    return () => {
      const item = isDelete ? 'customDialogDel' : 'customDialogAdd';
      self.setState({ [item]: true });
    };
  }

  handleSubmit = (cat, selectedOp, deleteOp = false) => {
    const self = this;
    return () => {
      const {
        namespace, cmEditorValAdd, cmEditorValDel,
      } = self.state;
      const { adapter } = self.props;
      const filteredOp = adapter.ops.filter(({ key }) => key === selectedOp);
      if (selectedOp === '' || typeof filteredOp === 'undefined' || filteredOp.length === 0) {
        self.setState({ selectionError: true });
        return;
      }
      if (deleteOp) {
        if (selectedOp === 'custom' && (cmEditorValDel === '' || self.cmEditorDel.state.lint.marked.length > 0)) {
          self.setState({ cmEditorValDelError: true, selectionError: true });
          return;
        }
      } else if (selectedOp === 'custom' && (cmEditorValAdd === '' || self.cmEditorAdd.state.lint.marked.length > 0)) {
        self.setState({ cmEditorValAddError: true, selectionError: true });
        return;
      }
      if (namespace === '') {
        self.setState({ namespaceError: true });
        return;
      }
      self.submitOp(cat, selectedOp, deleteOp);
    };
  }

  submitOp = (cat, selectedOp, deleteOp = false) => {
    const {
      namespace, cmEditorValAdd, cmEditorValDel, menuState,
    } = this.state;
    const { adapter } = this.props;
    // const fileInput = document.querySelector('#k8sfile') ;

    const data = {
      adapter: adapter.adapter_location,
      query: selectedOp,
      namespace,
      customBody: deleteOp ? cmEditorValDel : cmEditorValAdd,
      deleteOp: deleteOp ? 'on' : '',
    };

    const params = Object.keys(data).map((key) => `${encodeURIComponent(key)}=${encodeURIComponent(data[key])}`).join('&');
    this.props.updateProgress({ showProgress: true });
    const self = this;
    dataFetch('/api/mesh/ops', {
      credentials: 'same-origin',
      method: 'POST',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8',
      },
      body: params,
    }, (result) => {
      self.props.updateProgress({ showProgress: false });
      menuState[cat][deleteOp ? 'delete' : 'add'] = false;
      const dlg = deleteOp ? 'customDialogDel' : 'customDialogAdd';
      self.setState({ menuState, [dlg]: false });

      if (typeof result !== 'undefined') {
        self.props.enqueueSnackbar('Operation executing...', {
          variant: 'info',
          autoHideDuration: 2000,
          action: (key) => (
            <IconButton
              key="close"
              aria-label="Close"
              color="inherit"
              onClick={() => self.props.closeSnackbar(key)}
            >
              <CloseIcon />
            </IconButton>
          ),
        });
      }
    }, self.handleError(cat, deleteOp));
  }

  handleAdapterClick = (adapterLoc) => () => {
    this.props.updateProgress({ showProgress: true });
    const self = this;
    dataFetch(`/api/mesh/adapter/ping?adapter=${encodeURIComponent(adapterLoc)}`, {
      credentials: 'same-origin',
      credentials: 'include',
    }, (result) => {
      this.props.updateProgress({ showProgress: false });
      if (typeof result !== 'undefined') {
        this.props.enqueueSnackbar('Adapter successfully pinged!', {
          variant: 'success',
          autoHideDuration: 2000,
          action: (key) => (
            <IconButton
              key="close"
              aria-label="Close"
              color="inherit"
              onClick={() => self.props.closeSnackbar(key)}
            >
              <CloseIcon />
            </IconButton>
          ),
        });
      }
    }, self.handleError('Could not ping adapter.'));
  }

  handleSMIClick = (adapterLoc) => () => {
    this.props.updateProgress({ showProgress: true });
    const self = this;
    dataFetch(`/api/mesh/adapter/ping?adapter=${encodeURIComponent(adapterLoc)}`, {
      credentials: 'same-origin',
      credentials: 'include',
    }, (result) => {
      this.props.updateProgress({ showProgress: false });
      if (typeof result !== 'undefined') {
        self.setState({ ['customDialogSMI']: true })
      }
    }, self.handleError('Could not ping adapter.'));
  }

  handleError = (cat, deleteOp) => {
    const self = this;
    return (error) => {
      const { menuState } = self.state;
      menuState[cat][deleteOp ? 'delete' : 'add'] = false;
      const dlg = deleteOp ? 'customDialogDel' : 'customDialogAdd';
      self.setState({ menuState, [dlg]: false });

      self.props.updateProgress({ showProgress: false });
      self.props.enqueueSnackbar(`Operation submission failed: ${error}`, {
        variant: 'error',
        action: (key) => (
          <IconButton
            key="close"
            aria-label="Close"
            color="inherit"
            onClick={() => self.props.closeSnackbar(key)}
          >
            <CloseIcon />
          </IconButton>
        ),
        autoHideDuration: 8000,
      });
    };
  }

  handleExpandClick() {
    // setExpanded(!expanded);
  }

  generateMenu(cat, isDelete, selectedAdapterOps) {
    const { menuState } = this.state;
    const ele = !isDelete ? this.addIconEles[cat] : this.delIconEles[cat];
    return (
      <Menu
        id="long-menu"
        anchorEl={ele}
        keepMounted
        open={menuState[cat][isDelete ? 'delete' : 'add']}
        onClose={this.addDelHandleClick(cat, isDelete)}
      >
        {selectedAdapterOps.map(({ key, value }) => (
          <MenuItem key={`${key}_${new Date().getTime()}`} onClick={this.handleSubmit(cat, key, isDelete)}>
            {value}
          </MenuItem>
        ))}
      </Menu>
    );
  }

  handleOpen = () => {
    setOpen(true);
  }

  handleClose = () => {
    setOpen(false);
  }


  generateSMIResult() {
    const {
      customDialogSMI,
    } = this.state;

    const {
      smi_result,
    } = this.props;

    const columns = ["Test", "SMI Version", "Service Mesh", "Service Mesh Version", "SMI Specification", "Capability", "Test Status"];

    var data = [
      ["TA-01", "v1alpha3", "Linkerd", "edge-20.7.5", "Traffic Access", "Full", "Passed"],
      ["TA-02", "v1alpha3", "Linkerd", "edge-20.7.5", "Traffic Access", "Full", "Failed"],
      ["TM-01", "v1alpha3", "Linkerd", "edge-20.7.5", "Traffic Metrics", "Half", "Passed"],
      ["TM-02", "v1alpha3", "Linkerd", "edge-20.7.5", "Traffic Metrics", "None", "Passed"],
      ["TM-03", "v1alpha3", "Maesh", "v1.3.2", "Traffic Metrics", "None", "Failed"],
      ["TM-04", "v1alpha3", "Maesh", "v1.3.2", "Traffic Metrics", "Full", "Passed"],
    ];

    if(smi_result.details.results){
      data = smi_result.details.results.map((val) => {
        return [val.name, val.time, val.assertions,];
      });
    }

    return (
      <Dialog
        onClose={this.handleSMIClose()}
        aria-labelledby="adapter-dialog-title"
        open={customDialogSMI}
        fullWidth
        maxWidth="md"
      >
        <MUIDataTable
          title={"SMI Conformance Result"}
          data={data}
          columns={columns}
        />
      </Dialog>
    );
  }

  generateYAMLEditor(cat, isDelete) {
    const { adapter } = this.props;
    const {
      customDialogAdd, customDialogDel, namespace, namespaceError, cmEditorValAdd, cmEditorValDel,
    } = this.state;
    const self = this;
    return (
      <Dialog
        onClose={this.handleModalClose(isDelete)}
        aria-labelledby="adapter-dialog-title"
        open={isDelete ? customDialogDel : customDialogAdd}
        fullWidth
        maxWidth="md"
      >
        <DialogTitle id="adapter-dialog-title" onClose={this.handleModalClose(isDelete)}>
          {adapter.name}
          {' '}
          Adapter - Custom YAML
          {isDelete ? '(delete)' : ''}
        </DialogTitle>
        <Divider variant="fullWidth" light />
        <DialogContent>
          <Grid container spacing={5}>
            <Grid item xs={12}>
              <TextField
                required
                id="namespace"
                name="namespace"
                label="Namespace"
                fullWidth
                value={namespace}
                error={namespaceError}
                margin="normal"
                variant="outlined"
                onChange={this.handleChange('namespace')}
              />
            </Grid>
            <Grid item xs={12}>
              <CodeMirror
                editorDidMount={(editor) => {
                  if (isDelete) {
                    self.cmEditorDel = editor;
                  } else {
                    self.cmEditorAdd = editor;
                  }
                }}
                value={isDelete ? cmEditorValDel : cmEditorValAdd}
                options={{
                  theme: 'material',
                  lineNumbers: true,
                  lineWrapping: true,
                  gutters: ['CodeMirror-lint-markers'],
                  lint: true,
                  mode: 'text/x-yaml',
                }}
                onBeforeChange={(editor, data, value) => {
                  if (isDelete) {
                    self.setState({ cmEditorValDel: value });
                  } else {
                    self.setState({ cmEditorValAdd: value });
                  }
                  if (isDelete) {
                    if (value !== '' && self.cmEditorDel.state.lint.marked.length === 0) {
                      self.setState({ selectionError: false, cmEditorValDelError: false });
                    }
                  } else if (value !== '' && self.cmEditorAdd.state.lint.marked.length === 0) {
                    self.setState({ selectionError: false, cmEditorValAddError: false });
                  }
                }}
              />
            </Grid>
          </Grid>
        </DialogContent>
        <Divider variant="fullWidth" light />
        <DialogActions>
          <IconButton aria-label="Apply" color="primary" onClick={this.handleSubmit(cat, 'custom', isDelete)}>
            {/* <FontAwesomeIcon icon={faArrowRight} transform="shrink-4" fixedWidth /> */}
            {!isDelete && <PlayIcon />}
            {isDelete && <DeleteIcon />}
          </IconButton>

        </DialogActions>
      </Dialog>
    );
  }


  addDelHandleClick = (cat, isDelete) => {
    const self = this;
    return () => {
      const { menuState, customDialogAdd, customDialogDel } = self.state;
      menuState[cat][isDelete ? 'delete' : 'add'] = !menuState[cat][isDelete ? 'delete' : 'add'];

      const dlg = isDelete ? 'customDialogDel' : 'customDialogAdd';
      let dlgv = isDelete ? customDialogDel : customDialogAdd;
      if (cat === 4) {
        dlgv = !dlgv;
      }
      self.setState({ menuState, [dlg]: dlgv });
    };
  }

  generateCardForCategory(cat) {
    if (typeof cat === 'undefined') {
      cat = 0;
    }
    const { classes, adapter } = this.props;
    // const expanded = false;

    const selectedAdapterOps = adapter && adapter.ops ? adapter.ops.filter(({ category }) => typeof category === 'undefined' && cat === 0 || category === cat) : [];
    let content;
    let description;
    switch (cat) {
      case 0:
        content = 'Manage Service Mesh Lifecycle';
        description = 'Deploy a service mesh or SMI adapter on your cluster.';
        break;

      case 1:
        content = 'Manage Sample Application Lifecycle';
        description = 'Deploy sample applications on/off the service mesh.';
        break;

      case 2:
        content = 'Apply Service Mesh Configuration';
        description = 'Configure your service mesh using some pre-defined options.';
        break;

      case 3:
        content = 'Validate Service Mesh Configuration';
        description = 'Validate your service mesh configuration against best practices.';
        break;

      case 4:
        content = 'Apply Custom Configuration';
        description = 'Customize the configuration of your service mesh.';
        break;
    }
    return (
      <Card className={classes.card}>
        <CardHeader
          title={content}
          subheader={description}
        />
        <CardActions disableSpacing>
          <IconButton aria-label="install" ref={(ch) => this.addIconEles[cat] = ch} onClick={this.addDelHandleClick(cat, false)}>
            {cat !== 4 ? <AddIcon /> : <PlayIcon />}
          </IconButton>
          {cat !== 4 && this.generateMenu(cat, false, selectedAdapterOps)}
          {cat === 4 && this.generateYAMLEditor(cat, false)}
          {cat !== 3 && (
            <div className={classes.fileLabel}>
              <IconButton aria-label="delete" ref={(ch) => this.delIconEles[cat] = ch} className={classes.deleteRight} onClick={this.addDelHandleClick(cat, true)}>
                <DeleteIcon />
              </IconButton>
              {cat !== 4 && this.generateMenu(cat, true, selectedAdapterOps)}
              {cat === 4 && this.generateYAMLEditor(cat, true)}
            </div>
          )}
        </CardActions>
      </Card>
    );
  }

  render() {
    const {
      classes, adapter,
    } = this.props;
    const {
      namespace,
      namespaceError,
    } = this.state;

    let adapterName = (adapter.name).split(" ").join("").toLowerCase();
    let imageSrc = "/static/img/" + adapterName + ".svg";
    let adapterChip = (
      <Chip
        label={adapter.adapter_location}
        onClick={this.handleAdapterClick(adapter.adapter_location)}
        icon={<img src={imageSrc} className={classes.icon} />}
        className={classes.chip}
        variant="outlined"
      />
    );

    let imageSMISrc = "/static/img/smi.png";
    let smiChip = (
      <React.Fragment>
        <Chip
          label="View SMI Conformance results"
          onClick={this.handleSMIClick(adapter.adapter_location)}
          icon={<img src={imageSMISrc} className={classes.icon} />}
          className={classes.chip}
          variant="outlined"
        />
        {this.generateSMIResult()}
      </React.Fragment>
    );

    const filteredOps = [];
    if (adapter && adapter.ops && adapter.ops.length > 0) {
      adapter.ops.forEach(({ category }) => {
        if (typeof category === 'undefined') {
          category = 0;
        }
        if (filteredOps.indexOf(category) === -1) {
          filteredOps.push(category);
        }
      });
      filteredOps.sort();
    }

    return (
      <NoSsr>
        <React.Fragment>
          <div className={classes.chipGrid}>
            <Grid container spacing={3}>
              <Grid item xs={3}>
                {adapterChip}
              </Grid>
              <Grid item xs={3}>
              </Grid>
              <Grid item xs={3}>
              </Grid>
              <Grid item xs={3}>
                {smiChip}
              </Grid>
            </Grid>
          </div>
          <div className={classes.root}>
            <Grid container spacing={5}>
              <Grid item xs={12}>
                <TextField
                  required
                  id="namespace"
                  name="namespace"
                  label="Namespace"
                  fullWidth
                  value={namespace}
                  error={namespaceError}
                  margin="normal"
                  variant="outlined"
                  onChange={this.handleChange('namespace')}
                />
              </Grid>
              {filteredOps.map((val) => (
                <Grid item xs={12} md={4}>
                  {this.generateCardForCategory(val)}
                </Grid>
              ))}
            </Grid>
          </div>
        </React.Fragment>
      </NoSsr>
    );
  }
}

MesheryAdapterPlayComponent.propTypes = {
  classes: PropTypes.object.isRequired,
  adapter: PropTypes.object.isRequired,
};

const mapStateToProps = (state) => {
  const smi_result = state.get('smi_result').toJS();
  return { smi_result, };
};

const mapDispatchToProps = (dispatch) => ({
  updateProgress: bindActionCreators(updateProgress, dispatch),
});

export default withStyles(styles)(connect(
  mapStateToProps,
  mapDispatchToProps,
)(withRouter(withSnackbar(MesheryAdapterPlayComponent))));
