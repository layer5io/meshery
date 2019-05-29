import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import { NoSsr, Grid, ExpansionPanelDetails, Typography } from '@material-ui/core';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import LazyLoad from 'react-lazyload';
import GrafanaDateRangePicker from './GrafanaDateRangePicker';
import { ExpansionPanel, ExpansionPanelSummary } from './ExpansionPanels';
import GrafanaCustomChart from './GrafanaCustomChart';

const grafanaStyles = theme => ({
    root: {
      width: '100%',
    },
    column: {
      flexBasis: '33.33%',
    },
    heading: {
      fontSize: theme.typography.pxToRem(15),
    },
    secondaryHeading: {
      fontSize: theme.typography.pxToRem(15),
      color: theme.palette.text.secondary,
    },
    dateRangePicker: {
      display: 'flex',
      justifyContent: 'flex-end',
      marginRight: theme.spacing(1),
      marginBottom: theme.spacing(2),
    },
    iframe: {
      minHeight: theme.spacing(55),
      minWidth: theme.spacing(55),
    }
  });

class GrafanaCustomCharts extends Component {
  constructor(props) {
    super(props);

    const startDate = new Date();
    startDate.setMinutes(startDate.getMinutes() - 5);
    this.state = {
      startDate,
      from: 'now-5m',
      endDate: new Date(),
      to: 'now',
      liveTail: true,
      refresh: '10s',
    }

  }
    updateDateRange = (from, startDate, to, endDate, liveTail, refresh) => {
      this.setState({from, startDate, to, endDate, liveTail, refresh});
    }

    genRandomNumberForKey = () => {
      return Math.floor((Math.random() * 1000) + 1);
    }
    
    render() {
        const {from, startDate, to, endDate, liveTail, refresh} = this.state;
        const { classes, boardPanelConfigs } = this.props;
        let {grafanaURL} = this.props;
        if (grafanaURL.endsWith('/')){
          grafanaURL = grafanaURL.substring(0, grafanaURL.length - 1);
        }
        return (
              <NoSsr>
              <React.Fragment>
              <div className={classes.root}>
                <div className={classes.dateRangePicker}>
                  <GrafanaDateRangePicker from={from} startDate={startDate} to={to} endDate={endDate} liveTail={liveTail} 
                    refresh={refresh} updateDateRange={this.updateDateRange} />
                </div>
                {boardPanelConfigs.map((config, ind) => (
                  // <ExpansionPanel defaultExpanded={ind === 0?true:false}>
                  <ExpansionPanel square defaultExpanded={false}>
                    <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />}>
                      <div className={classes.column}>
                      <Typography variant="subtitle1" gutterBottom>{config.board.title}</Typography>
                      </div>
                      <div className={classes.column}>
                        <Typography variant="subtitle2">{config.templateVars && config.templateVars.length > 0?'Template variables: '+config.templateVars.join(' '):''}</Typography>
                      </div>
                    </ExpansionPanelSummary>
                    <ExpansionPanelDetails>
                        <Grid container spacing={5}>
                          {config.panels.map(panel => (
                            <Grid item xs={12} sm={6} className={classes.iframe}>
                              <GrafanaCustomChart
                                key={this.genRandomNumberForKey()}
                                board={config}
                                panel={panel}
                                grafanaURL={grafanaURL}
                                from={from} startDate={startDate} to={to} endDate={endDate} liveTail={liveTail} refresh={refresh}
                              />
                            </Grid>
                          ))}
                        </Grid>
                    </ExpansionPanelDetails>
                  </ExpansionPanel>
                ))}
              </div>
              </React.Fragment>
              </NoSsr>
            );
        }
}

GrafanaCustomCharts.propTypes = {
  classes: PropTypes.object.isRequired,
  grafanaURL: PropTypes.string.isRequired,
  boardPanelConfigs: PropTypes.array.isRequired,
};

export default withStyles(grafanaStyles)(GrafanaCustomCharts);