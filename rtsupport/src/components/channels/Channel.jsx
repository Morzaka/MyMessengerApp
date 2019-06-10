import React, {Component} from 'react';

class Channel extends Component {
    onClick(e){
        e.preventDefault();
        const {setChannel, chennel} = this.props;
        setChannel(chennel);
    }

    render() {
        const {channel} = this.props;
        return (
            <li>
                <a>
                    {channel.name}
                </a>
            </li>
        )
    }
}

Channel.propType = {
    channel: React.PropTypes.object.isRequired,
    setChannel: React.PropTypes.func.isRequired
};

export default Channel;