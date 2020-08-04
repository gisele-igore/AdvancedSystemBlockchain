import React from "react";
import {
  ListGroup,
  ListGroupItem,
  Row,
  Col,
  Form,
  FormInput,
  FormGroup,
  FormCheckbox,
  FormSelect,
  Button
} from "shards-react";

const CompleteFormExample = () => (
  <ListGroup flush>
    <ListGroupItem className="p-3">
      <Row>
        <Col>
          <Form>
            <Row form>
              <Col md="6" className="form-group">
                <label>Insurance Company</label>
                <FormInput
                  id="IDCompagnieAssurance"
                  type="text"
                  placeholder="Company Name"
                />
              </Col>

              <Col md="6">
                <label>Insurance Buyer</label>
                <FormInput
                  id="CodeAcheteurAssurance"
                  type="text"
                  placeholder="Buyer's Name"
                />
              </Col>
            </Row>

            <Row form>
              <Col md="6" className="form-group">
                <label>From</label>
                <FormInput id="DateDebut" type="date" placeholder="" />
              </Col>

              <Col md="6">
                <label>To</label>
                <FormInput id="DateDebut" type="date" placeholder="" />
              </Col>
            </Row>

            <FormGroup>
              <label>Contract File</label>
                <div className="custom-file mb-3">
                  <input type="file" className="custom-file-input" id="customFile2" />
                  <label className="custom-file-label" htmlFor="customFile2">
                    Choose file...
                  </label>
                </div>
            </FormGroup>

            <Row form>
              <Col md="6" className="form-group">
                <label>Company Signature</label>
                <FormInput id="SignatureCompagnie" />
              </Col>
              <Col md="6" className="form-group">
                <label>Buyer Signature</label>
                <FormInput id="SignatureAcheteur" />
              </Col> 
              <Col md="12" className="form-group">
                <FormCheckbox>
                  {/* eslint-disable-next-line */}I agree with your{" "}
                  <a href="#">Privacy Policy</a>.
                </FormCheckbox>
              </Col>
            </Row>
            <Button type="submit">Sell New Contract</Button>
          </Form>
        </Col>
      </Row>
    </ListGroupItem>
  </ListGroup>
);

export default CompleteFormExample;
