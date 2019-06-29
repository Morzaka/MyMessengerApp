import React, {Component} from 'react';
import User from './User.jsx'
import PropTypes from 'prop-types';

class UserList extends Component {
    render() {
        return (
            <ul className="list-group">{
                this.props.users.map(usr => {
                    return (
                        <User
                        user={usr}
                        key={usr.id}
                    />)
                })
            }</ul>
        )
    }
}

UserList.propTypes = {
    users: PropTypes.array.isRequired
};

export default UserList;