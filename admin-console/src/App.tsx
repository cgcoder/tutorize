import "normalize.css";
import "@blueprintjs/core/lib/css/blueprint.css";
// include blueprint-icons.css for icon font support
import "@blueprintjs/icons/lib/css/blueprint-icons.css";

import { Classes, Spinner, Button, Card, H3 } from "@blueprintjs/core";

function App() {
  return (
    <>
      <div>
        <Spinner className={Classes.SMALL} intent="primary" />
        <Button text="Click Me" />

        <Card>
          <H3>Adventure Awaits</H3>
          <p>
            Embark on an epic journey across uncharted lands. This card outlines
            your mission.
          </p>
          <Button intent="primary">Start Journey</Button>
        </Card>
      </div>
    </>
  );
}

export default App;
