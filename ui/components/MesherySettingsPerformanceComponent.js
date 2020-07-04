/* eslint-disable */ 
import React from 'react';
import PropTypes from 'prop-types';
import Button from '@material-ui/core/Button';
import { withStyles } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';
import {
  NoSsr, Tooltip, IconButton, CircularProgress, FormControl, RadioGroup, FormControlLabel, Radio,
} from '@material-ui/core';
import dataFetch from '../lib/data-fetch';
import TextField from '@material-ui/core/TextField';
import { withSnackbar } from 'notistack';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import CloseIcon from '@material-ui/icons/Close';
import { updateLoadTestPref, updateProgress } from '../lib/store';


const loadGenerators = [
  'fortio',
  'wrk2',
];

const styles = (theme) => ({
  root: {
    padding: theme.spacing(10),
  },
  buttons: {
    display: 'flex',
    justifyContent: 'flex-end',
  },
  button: {
    marginTop: theme.spacing(3),
    marginLeft: theme.spacing(1),
  },
  margin: {
    margin: theme.spacing(1),
  },
});

class MesherySettingsPerformanceComponent extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      qps: '',
      c: '',
      t: '',
      gen: '',
      blockRunTest: false,
      tError: false,
    };
  }

  handleChange = (name) => (event) => {
    if (name === 't' && (event.target.value.toLowerCase().endsWith('h')
      || event.target.value.toLowerCase().endsWith('m') || event.target.value.toLowerCase().endsWith('s'))) {
      this.setState({ tError: false });
    }
    this.setState({ [name]: event.target.value });
  };

  handleSubmit = () => {
    const {
      t
    } = this.state;

    let err = false; 
    let tNum = 0;
    try {
      tNum = parseInt(t.substring(0, t.length - 1));
    } catch (ex) {
      err = true;
    }

    if (t === '' || !(t.toLowerCase().endsWith('h')
      || t.toLowerCase().endsWith('m') || t.toLowerCase().endsWith('s')) || err || tNum <= 0) {
      this.setState({ tError: true });
      return;
    }

    this.submitLoadTest();
  }

  submitLoadTest = () => {
    const {
      qps, c, t, gen,
    } = this.state;

    const data = {
      qps,
      c,
      t,
      gen,
    };
    const params = Object.keys(data).map((key) => `${encodeURIComponent(key)}=${encodeURIComponent(data[key])}`).join('&');

    this.setState({ blockRunTest: true }); // to block the button
    this.props.updateProgress({ showProgress: true });
    const self = this;
    dataFetch('/api/load-test-prefs', {
      credentials: 'same-origin',
      method: 'POST',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8',
      },
      body: params,
    }, (result) => {
      this.props.updateProgress({ showProgress: false });
      if (typeof result !== 'undefined') {
        this.props.enqueueSnackbar('Preference was successfully updated!', {
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
        this.props.updateLoadTestPref({
          loadTestPref: {
            qps: self.state.qps,
            c: self.state.c,
            t: self.state.t,
            gen: self.state.gen,
          }
        });
        this.setState({ blockRunTest: false });
      }
    }, self.handleError('There was an error saving your preferences'));
  }

  componentDidMount() {
    this.getLoadTestPrefs();
  }

  getLoadTestPrefs = () => {
    const self = this;
    dataFetch('/api/load-test-prefs', {
      credentials: 'same-origin',
      method: 'GET',
      credentials: 'include',
    }, (result) => {
      if (typeof result !== 'undefined') {
        this.setState({               
            qps: result.loadTestPrefs.qps,
            c: result.loadTestPrefs.c,
            t: result.loadTestPrefs.t,
            gen: result.loadTestPrefs.gen,
          });
      }
    }, self.handleError('There was an error fetching your preferences'));
  }

  handleError = (msg) => {
    const self = this;
    return (error) => {
      self.setState({ blockRunTest: false });
      let finalMsg = msg;
      if (typeof error === 'string') {
        finalMsg = `${msg}: ${error}`;
      }
      self.props.enqueueSnackbar(finalMsg, {
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
        autoHideDuration: 4000,
      });
    };
  }

  render() {
    const { classes } = this.props;
    const {
      blockRunTest, qps, t, c, gen,
      tError,
    } = this.state;
    console.log(gen);

    return (

      <NoSsr>
        <React.Fragment>
          <div className={classes.root}>
            <label><strong>Performance Load Test Defaults</strong></label>
            <Grid container spacing={3}>
              <Grid item xs={12} sm={4}>
                <TextField
                  required
                  id="c"
                  name="c"
                  label="Concurrent requests"
                  type="number"
                  fullWidth
                  value={c}
                  inputProps={{ min: '0', step: '1' }}
                  margin="normal"
                  variant="outlined"
                  onChange={this.handleChange('c')}
                />
              </Grid>
              <Grid item xs={12} sm={4}>
                <TextField
                  required
                  id="qps"
                  name="qps"
                  label="Queries per second"
                  type="number"
                  fullWidth
                  value={qps}
                  inputProps={{ min: '0', step: '1' }}
                  margin="normal"
                  variant="outlined"
                  onChange={this.handleChange('qps')}
                />
              </Grid>
              <Grid item xs={12} sm={4}>
                <Tooltip title={"Please use 'h', 'm' or 's' suffix for hour, minute or second respectively."}>
                  <TextField
                    required
                    id="t"
                    name="t"
                    label="Duration"
                    fullWidth
                    value={t}
                    error={tError}
                    margin="normal"
                    variant="outlined"
                    onChange={this.handleChange('t')}
                  />
                </Tooltip>
              </Grid>
              <Grid item xs={12} sm={4}>
                <FormControl component="loadGenerator" className={classes.formControl}>
                  <label><strong>Default Load Generator</strong></label>
                  <RadioGroup aria-label="loadGenerator" name="loadGenerator" value={gen} onChange={this.handleChange('gen')} row>
                    {loadGenerators.map((lg) => (
                      <FormControlLabel value={lg} control={<Radio color="primary" />} label={lg} />
                    ))}
                  </RadioGroup>
                </FormControl>
              </Grid>
            </Grid>
              <div className={classes.buttons}>
                <Button
                  type="submit"
                  variant="contained"
                  color="primary"
                  size="large"
                  onClick={this.handleSubmit}
                  className={classes.button}
                  disabled={blockRunTest}
                >
                  {blockRunTest ? <CircularProgress size={30} /> : 'Submit Preferences'}
                </Button>
              </div>
          </div>
        </React.Fragment>
      </NoSsr>
    );
  }
}

MesherySettingsPerformanceComponent.propTypes = {
  classes: PropTypes.object.isRequired,
};

const mapDispatchToProps = (dispatch) => ({
  updateLoadTestPref: bindActionCreators(updateLoadTestPref, dispatch),
  updateProgress: bindActionCreators(updateProgress, dispatch),
});

const mapStateToProps = (state) => {

  const loadTestPref = state.get('loadTestPref').toJS();
  return {
    ...loadTestPref,
  };
};


export default withStyles(styles)(connect(
    mapStateToProps,
    mapDispatchToProps,
)(withSnackbar(MesherySettingsPerformanceComponent)));
