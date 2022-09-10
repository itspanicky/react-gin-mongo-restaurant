import 'bootstrap/dist/css/bootstrap.css'
import { Button, Card, Row, Col } from 'react-bootstrap'
import { OrderData, OrderId, ChangeOrderOrWaiter } from '../types'

export default function Order({
  orderData,
  setChangeWaiter,
  setChangeOrder,
  deleteSingleOrder,
}: Props) {
  return (
    <Card>
      <Row>
        <Col>Dish: {orderData && orderData.dish}</Col>
        <Col>Server: {orderData && orderData.server}</Col>
        <Col>Table: {orderData && orderData.table}</Col>
        <Col>Price: ${orderData && orderData.price}</Col>
        <Col><Button onClick={() => deleteSingleOrder(orderData._id)}>delete order</Button></Col>
        <Col><Button onClick={() => changeWaiter()}>change waiter</Button></Col>
        <Col><Button onClick={() => changeOrder()}>change order</Button></Col>
      </Row>
    </Card>
  )
  function changeWaiter() {
    setChangeWaiter({
      change: true,
      id: orderData._id
    })
  }
  function changeOrder() {
    setChangeOrder({
      change: true,
      id: orderData._id
    })
  }
}

type Props = {
  orderData: OrderData & OrderId
  setChangeWaiter: ({ change, id }: ChangeOrderOrWaiter) => void
  setChangeOrder: ({ change, id }: ChangeOrderOrWaiter) => void
  deleteSingleOrder: (id: number) => void
}