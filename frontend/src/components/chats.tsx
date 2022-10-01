function Chats(props: { chatHistory: Array<any> }) {
  const messages = props.chatHistory.map((msg, index) => (
    <h2 key={index}>{msg}</h2>
  ));

  return <div className="Chats">{messages}</div>;
}

export { Chats };
