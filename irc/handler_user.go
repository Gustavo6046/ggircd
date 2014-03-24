package irc

type commandMap map[string]func(State, *User, Connection, Message) Handler

// UserHandler is a handler that handles messages coming from a user connection
// that has successfully associated with the client.
type UserHandler struct {
  state    chan State
  nick     string
  commands commandMap
}

func NewUserHandler(state chan State, nick string) Handler {
  handler := &UserHandler{
    state: state,
    nick:  nick,
  }
  handler.commands = commandMap{
    CmdJoin.Command: handler.handleCmdJoin,
    CmdPing.Command: handler.handleCmdPing,
    CmdQuit.Command: handler.handleCmdQuit,
  }
  return handler
}

func (h *UserHandler) Closed(conn Connection) {
  conn.Kill()
}

func (h *UserHandler) Handle(conn Connection, msg Message) Handler {
  state := <-h.state
  defer func() { h.state <- state }()

  command := h.commands[msg.Command]
  if command == nil {
    return h
  }
  return command(state, state.GetUser(h.nick), conn, msg)
}
