import React, { useState } from 'react';
import {
  Card,
  Button,
  FormGroup,
  InputGroup,
  Divider,
  Callout,
  H2,
  Intent,
} from '@blueprintjs/core';
import { Link } from 'react-router';
import './LoginPage.css';
import Page from '../../components/Page';

const LoginPage: React.FC = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [isLoading, setIsLoading] = useState(false);

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    setError('');

    try {
      // Implement your authentication logic here
      console.log('Logging in with:', username, password);
      // Example: await authService.login(username, password);

      // Clear form after successful login
      setUsername('');
      setPassword('');
    } catch (err) {
      setError('Invalid username or password');
    } finally {
      setIsLoading(false);
    }
  };

  const handleGoogleLogin = () => {
    // Implement Google authentication logic
    console.log('Logging in with Google');
    // Example: window.location.href = '/api/auth/google';
  };

  return (
    <Page>
      <div className="login-container">
        <Card elevation={2} className="login-card">
          <H2 className="login-title">Tutorize Admin</H2>

          {error && <Callout intent={Intent.DANGER} title="Login Failed">{error}</Callout>}

          <form onSubmit={handleLogin}>
            <FormGroup label="Username" labelFor="username">
              <InputGroup
                id="username"
                placeholder="Enter your username"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                size="medium"
                required
              />
            </FormGroup>

            <FormGroup label="Password" labelFor="password">
              <InputGroup
                id="password"
                placeholder="Enter your password"
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                size="medium"
                required
              />
            </FormGroup>

            <Button
              type="submit"
              intent={Intent.PRIMARY}
              text="Log In"
              fill
              large
              loading={isLoading}
            />
          </form>
          <Divider style={{ margin: '25px' }} />
          <Button
            text="Sign in with Google"
            fill
            onClick={handleGoogleLogin}
            className="google-button"
          />

          <div className="register-link">
            <p>
              Don't have an account?{' '}
              <Link to="/register">Register here</Link>
            </p>
            <p>
              Forgot your password?{' '}
              <Link to="/reset_password">Reset here</Link>
            </p>
          </div>


        </Card>
      </div>
    </Page>
  );
};

export default LoginPage;
