import React, {Component} from 'react';
import PropTypes from 'prop-types';
import fecha from 'fecha';

class Message extends Component {
    render() {
        let {message} = this.props;
        let createdAt = fecha.format(message.createdAt, 'HH:mm:ss MM/DD/YY');
        return (
            <li className='message card'>
                <div className='author'>
                    <strong>{message.author}</strong>
                    <i className='timestamp card-subtitle mb-2 text-muted'>{createdAt}</i>
                </div>
                <div className='body card-text'>{message.body}</div>
            </li>
        )
    }
}

Message.propTypes = {
    message: PropTypes.object.isRequired
};

export default Message