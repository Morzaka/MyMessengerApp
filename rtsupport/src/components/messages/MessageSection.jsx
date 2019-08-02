import React, {Component} from 'react';
import MessageList from './MessageList.jsx';
import MessageForm from './MessageForm.jsx';
import PropTypes from 'prop-types';

class MessageSection extends Component {
    render() {
        let {activeChannel} = this.props;
        return (
            <div className='messages-container card bg-light'>
                <div className='card-header'><strong>{activeChannel.name || 'Select A Channel'}</strong></div>
                <div className='messages'>
                    <MessageList {...this.props} />
                    <MessageForm {...this.props} />
                </div>
            </div>
        )
    }
}

MessageSection.propTypes = {
    messages: PropTypes.array.isRequired,
    activeChannel: PropTypes.object.isRequired,
    addMessage: PropTypes.func.isRequired
};

export default MessageSection;