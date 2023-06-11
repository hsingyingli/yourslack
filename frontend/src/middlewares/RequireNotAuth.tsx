import {useLocation, Navigate, Outlet} from 'react-router-dom';
import useAuth from '../hooks/useAuth';

const RequireNotAuth = () => {
  const {user} = useAuth();
  const location = useLocation();
  return !user ? (
    <Outlet />
  ) : (
    <Navigate to={location.state?.from || '/'} replace />
  );
};

export {
  RequireNotAuth
}
