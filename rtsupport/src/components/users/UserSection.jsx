import React, {Component} from 'react';
import UserList from './UserList.jsx';
import UserForm from './UserForm.jsx';
import PropTypes from 'prop-types';

class UserSection extends Component {
    render() {
        return (
            <div className='support card bg-light'>
                <div className='card-header text-white bg-primary'>
                    <strong>Users</strong>
                </div>
                <div className='card-body users'>
                    <UserList {...this.props}/>
                    <UserForm {...this.props} />
                </div>
            </div>
        )
    }
}

UserSection.propTypes = {
    users: PropTypes.array.isRequired,
    setUserName: PropTypes.func.isRequired
};

export default UserSection;