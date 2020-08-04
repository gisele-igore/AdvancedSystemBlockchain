import React from "react";
import { Redirect } from "react-router-dom";
import { Container, Row, Col, Card, CardBody, Button } from "shards-react";
/* import NormalButtons from "../components/components-overview/NormalButtons"; */
import ComponentsOverview from "./ComponentsOverview"; 

import PageTitle from "../components/common/PageTitle";

const Tables = () => (
  <Container fluid className="main-content-container px-4">
    {/* Page Header */}
    <Row noGutters className="page-header py-4">
      <PageTitle sm="4" title="My Contracts" subtitle="" className="text-sm-left" />
    </Row>

    {/* Default Light Table */}
    <Row>
      <Col>
        <Card small className="mb-4">
          {/* <CardHeader className="border-bottom">
            <h6 className="m-0">Active Users</h6>
          </CardHeader> */}
          <CardBody className="p-0 pb-3">
            <table className="table mb-0">
              <thead className="bg-light">
                <tr>
                  <th scope="col" className="border-0">
                    #
                  </th>
                  <th scope="col" className="border-0">
                    Insurance Buyer
                  </th>
                  <th scope="col" className="border-0">
                    Insurance Company
                  </th>
                  <th scope="col" className="border-0">
                    From
                  </th>
                  <th scope="col" className="border-0">
                    To
                  </th>
                  <th scope="col" className="border-0">
                    Buyer Signature
                  </th>
                  <th scope="col" className="border-0">
                    Insurer Signature
                  </th>
                  <th scope="col" className="border-0">
                    Actions
                  </th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td>1</td>
                  <td>Ali</td>
                  <td>Kerry</td>
                  <td>Russian Federation</td>
                  <td>Gdańsk</td>
                  <td>107-0339</td>
                  <td>107-0339</td>
                  <td>
                    <Button theme="primary" className="mb-2 mr-1">Display</Button>
                    <Button theme="primary" className="mb-2 mr-1" href={ComponentsOverview}>Add New</Button>
                  </td>
                </tr>
                {/* <tr>
                  <td>2</td>
                  <td>Clark</td>
                  <td>Angela</td>
                  <td>Estonia</td>
                  <td>Borghetto di Vara</td>
                  <td>1-660-850-1647</td>
                </tr>
                <tr>
                  <td>3</td>
                  <td>Jerry</td>
                  <td>Nathan</td>
                  <td>Cyprus</td>
                  <td>Braunau am Inn</td>
                  <td>214-4225</td>
                </tr>
                <tr>
                  <td>4</td>
                  <td>Colt</td>
                  <td>Angela</td>
                  <td>Liberia</td>
                  <td>Bad Hersfeld</td>
                  <td>1-848-473-7416</td>
                </tr> */}
              </tbody>
            </table>
          </CardBody>
        </Card>
      </Col>
    </Row>
  </Container>
);

export default Tables;
