// ---------------------------------------------
// Chatkit Code
// ---------------------------------------------


const tokenProvider = new Chatkit.TokenProvider({
	url: `https://us1.pusherplatform.io/services/chatkit_token_provider/v1/05f46048-3763-4482-9cfe-51ff327c3f29/token`
})

const chatManager = new Chatkit.ChatManager({
	instanceLocator: "v1:us1:05f46048-3763-4482-9cfe-51ff327c3f29",
	userId: "admin",
	tokenProvider
})

 chatManager.connect().then(currentUser => {
	// // // setUser(user)
	currentUser
	.subscribeToRoom({
	 roomId: currentUser.rooms[0].id,
			hooks: { onNewMessage: addMessage },
			messageLimit: 10
		 })
		 // .then(setRoom)
})
function addMessage(message){
	console.log(message)
}
function sendMessage(textContent){
	currentUser.sendMessage({
  text: textContent,
  roomId: currentUser.rooms[0].id
});

// chatManager
  // .connect()
  // .then(currentUser => {
    // console.log("Connected as user ", currentUser);
  // })
  // .catch(error => {
    // console.error("error:", error);
  // });


// // ---------------------------------------------
// // Application Code
// // ---------------------------------------------


// const { app, h } = hyperapp
// /* @jsx h */

// const state = {
	// user: {
		// avatarURL:
			// "data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7"
	// },
	// room: {},
	// messages: []
// }

// const actions = {
	// setUser: user => ({ user }),
	// setRoom: room => ({ room }),
	// addMessage: payload => ({ messages }) => ({
		// messages: [payload, ...messages]
	// })
// }

// const { addMessage, setUser, setRoom } = app(
	// state,
	// actions,
	// view,
	// document.body
// )


// // ---------------------------------------------
// // View Code
// // ---------------------------------------------


// const UserHeader = ({ user }) => (
	// <header>
		// <img src={user.avatarURL} />
		// <p>{user.name}</p>
	// </header>
// )

// const MessageList = ({ user, messages }) => (
	// <div>
		// {messages.map(message => (
			// <message- class={message.sender.id === user.id && "own"}>
				// <img src={message.sender.avatarURL} />
				// <p>{message.text}</p>
			// </message->
		// ))}
	// </div>
// )

// const MessageInput = ({ user, room }) => (
	// <form
		// onsubmit={e => {
			// e.preventDefault()
			// user
				// .sendMessage({
					// text: e.target.elements[0].value,
					// roomId: room.id
				// })
				// .then(() => {
					// e.target.elements[0].value = ""
				// })
		// }}
	// >
		// <input placeholder="Type a message.." />
		// <button>
			// <svg>
				// <use href="#send" />
			// </svg>
		// </button>
	// </form>
// )

// function view(state, actions) {
	// return <main>
		// <UserHeader user={state.user} />
		// <MessageList user={state.user} messages={state.messages} />
		// <MessageInput user={state.user} room={state.room} />
	// </main>
// }
