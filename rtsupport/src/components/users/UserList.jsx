import React, {Component} from 'react';
import User from './User.jsx'
import PropTypes from 'prop-types';

class UserList extends Component {
    render() {
        return (
            <ul className="list-group">{
                this.props.users.map(user => {
                    return (
                        <User
                        user={user}
                        key={user.id}
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