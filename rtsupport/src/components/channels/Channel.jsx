import React, {Component} from 'react';
import PropTypes from 'prop-types';

class Channel extends Component {
  onClick(e) {
    e.preventDefault();
    const {setChannel, channel} = this.props;
    setChannel(channel);
  }

  render() {
    const {channel, activeChannel} = this.props;
    const active = activeChannel === channel ? 'active' : '';
    return (
      <li className={active}>
        <a onClick={this.onClick.bind(this)}>
          {channel.name}
        </a>
      </li>
    )
  }
}

Channel.propTypes = {
  channel: PropTypes.object.isRequired,
  setChannel: PropTypes.func.isRequired,
  activeChannel: PropTypes.object
};

export default Channel