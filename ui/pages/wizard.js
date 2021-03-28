import React from 'react';
import { Paper, withStyles, NoSsr } from '@material-ui/core';
import { connect } from "react-redux";
import { bindActionCreators } from 'redux'
import Head from 'next/head';

import { updatepagepath, updatepagetitle } from "../lib/store";
import { getPath } from "../lib/path";
import ConfigurationWizard from '../components/configuration-wizard/ConfigurationWizard'

const styles = {
  paper: {
    maxWidth: '90%',
    margin: 'auto',
    overflow: 'hidden',
  }
}

class Wizard extends React.Component {
  componentDidMount () {
    console.log(`path: ${getPath()}`);
    this.props.updatepagepath({ path: getPath() });
    this.props.updatepagetitle({ title: 'Configuration Wizard' });
  }

  render () {
    return (
      <NoSsr>
        <Head>
          <title>Configuration Wizard | Meshery</title>
        </Head>
        <Paper className={this.props.classes.paper}>
          <ConfigurationWizard/>
        </Paper>
      </NoSsr>
    );
  }
}

const mapDispatchToProps = dispatch => ({
  updatepagepath: bindActionCreators(updatepagepath, dispatch),
  updatepagetitle: bindActionCreators(updatepagetitle, dispatch),
})

export default withStyles(styles)(connect(
  null,
  mapDispatchToProps
)(Wizard));