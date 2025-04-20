import {
    Outlet,
} from "react-router";
import LinkButton from "./../../components/LinkButton";

export default function ConsolePage() {
    return (
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
                        <LinkButton to="/console" icon="home" text="Dashboard" />
                    </li>
                    <li>
                        <LinkButton to="/console/profile" icon="properties" text="Profile" />
                    </li>
                    <li>
                        <LinkButton to="/console/settings" icon="settings" text="Settings" />
                    </li>
                </ul>
            </nav>
            <div style={{ display: "flex", width: "100%"}}>
                <Outlet />
            </div>
        </div>);
}