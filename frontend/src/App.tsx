import {BrowserRouter as Router, Routes, Route} from 'react-router-dom';
import HomePage from './pages/HomePage';
import SignInPage from './pages/SignInPage';
import {RequireAuth} from './middlewares'

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/signin" element={<SignInPage />} />
        {/*   Protect Routes   */}
        <Route element={<RequireAuth />}>
          <Route path="/" element={<HomePage />} />
        </Route>
        {/*   Protect Routes   */}
      </Routes>
    </Router>
  );
}

export default App;
