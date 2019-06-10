import React, {Component} from 'react';
import ChannelSection from './channels/ChannelSection.jsx'

class App extends Component {
    constructor(props) {
        super(props);
        this.state = {
            channels: []
        };
    }
    addChannel(name){
        let {channels} = this.state;
        channels.push({id: channels.length, name});
        this.setState({channels});
        // TODO: send to server
    }
    setChannel(activeChannel){
        this.setState({activeChannel});
        // TODO: get channels messages
    }
    render() {
        return (
            <div>
                <ChannelSection
                    channels={this.state.channels}
                />
            </div>
        );
    }
}

export default App;