import "normalize.css";
import "@blueprintjs/core/lib/css/blueprint.css";
import "@blueprintjs/icons/lib/css/blueprint-icons.css";

import {
  Classes,
  Spinner,
  Button,
  Card,
  H3,
  Navbar,
  Alignment,
} from "@blueprintjs/core";
import {
  BrowserRouter as Router,
  Routes,
  Route,
  Link,
  Outlet,
} from "react-router";
import LinkButton from "./components/LinkButton";

function Dashboard() {
  return (
    <Card>
      <H3>Dashboard</H3>
      <p>Welcome to the dashboard!</p>
    </Card>
  );
}

function Profile() {
  return (
    <Card>
      <H3>Profile</H3>
      <p>Manage your profile information here.</p>

      <ul style={{ listStyle: "none", padding: 0 }}>
        <li>
          <Link to="create">Create Profile</Link>
        </li>
        <li>
          <Link to="edit">Edit Profile</Link>
        </li>
      </ul>
      <Outlet />
    </Card>
  );
}

function Settings() {
  return (
    <Card>
      <H3>Settings</H3>
      <p>Adjust your application settings here.</p>
    </Card>
  );
}

function App() {
  return (
    <Router>
      <div style={{ display: "flex", height: "100vh" }}>
        <nav
          style={{
            background: "#f4f4f4",
            padding: "5px",
            boxShadow: "2px 0 5px rgba(0,0,0,0.1)",
          }}
        >
          <ul style={{ listStyle: "none", padding: 0 }}>
            <li>
              <LinkButton to="/" icon="home" text="Dashboard" />
            </li>
            <li>
              <LinkButton to="/profile" icon="properties" text="Profile" />
            </li>
            <li>
              <LinkButton to="/settings" icon="settings" text="Settings" />
            </li>
          </ul>
        </nav>

        {/* Main Content Area */}
        <main style={{ flex: 1, padding: "20px" }}>
          <Routes>
            <Route path="/" element={<Dashboard />} />
            <Route path="/profile" element={<Profile />}>
              <Route path="create" element={<span>Create</span>} />
              <Route path="edit" element={<span>Edit</span>} />
            </Route>
            <Route path="/settings" element={<Settings />} />
          </Routes>
        </main>
      </div>
    </Router>
  );
}

export default App;
