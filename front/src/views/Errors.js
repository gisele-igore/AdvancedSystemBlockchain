import React from "react";
import { Container, Button, Col } from "shards-react";

import Editor from "../components/add-new-post/Editor";

const Errors = () => (
  <Container fluid className="main-content-container px-4 pb-4">
    {/* Editor */}
      <Col lg="9" md="12">
        <Editor />
        <Button pill>&larr; Save </Button>
      </Col>
    {/* <div className="error">
      <div className="error__content">
        <h2>500</h2>
        <h3>Something went wrong!</h3>
        <p>There was a problem on our end. Please try again later.</p>
        <Button pill>&larr; Go Back</Button>
      </div>
    </div> */}
  </Container>
);

export default Errors;

