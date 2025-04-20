import { Card, H3 } from "@blueprintjs/core";
import Page from "../../../components/Page";

export default function DashboardPage() {
    return (
        <Page>
            <Card style={{width: "100%"}}>
                <H3>Dashboard</H3>
                <p>Welcome to the dashboard!</p>
            </Card>
        </Page>
    );
}