package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForNicaragua(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<defs><linearGradient id="emojioneFlagForNicaragua0" x1="31.902" x2="32.494" y1="35.512" y2="33.945" gradientTransform="matrix(1 0 0 -1 0 66)" gradientUnits="userSpaceOnUse"><stop stop-color="#ff2a2a"/><stop offset="1" stop-color="red"/></linearGradient></defs><path fill="#428bc1" d="M32 2C18.9 2 7.8 10.4 3.7 22h56.6C56.2 10.4 45.1 2 32 2zm0 60c13.1 0 24.2-8.3 28.3-20H3.7C7.8 53.7 18.9 62 32 62z"/><path fill="#f9f9f9" d="M3.7 22C2.6 25.1 2 28.5 2 32s.6 6.9 1.7 10h56.6c1.1-3.1 1.7-6.5 1.7-10s-.6-6.9-1.7-10H3.7z"/><path fill="#42ade2" d="m32 25.3l-6.3 10.1h12.6z"/><path fill="#428bc1" d="m28.9 30.4l-3.2 5h12.6l-3.2-5z"/><path fill="url(#emojioneFlagForNicaragua0)" d="M33 31v-.4c0-.3-.2-.7-.5-1.1c-.2-.4-.5-.5-1-.2c-.2.1-.3.2-.4.5v.3s.2.1.3-.1c.1 0 .1-.1.1-.1c-.1 0-.1.1-.1.1c-.1.1-.2.2-.3.4c-.1.2-.1.4-.1.5c0 0 .1.2.3.2c.2.1.3.2.4.3c.1.1.1.2.1.4c0 .1 0 .3.1.3s.2-.1.3-.2c0-.2.1-.3.1-.4c.1-.1.1-.1.3-.1c.1 0 .2 0 .2-.1v-.1c.2 0 .1-.1.2-.2"/><path fill="#83bf4f" d="M37.6 34.4c-1-.1-1.9-1-2.1-1.9c0 0-.1-.1-.2-.1c-.2 0-.2.1-.2.1c-.1.3-.3.7-.6 1c-.3-.3-.5-.7-.6-1c0 0-.1-.1-.2-.1c-.2 0-.2.1-.2.1c-.1.3-.3.7-.6 1c-.3-.3-.5-.7-.6-1c0 0-.1-.1-.2-.1s-.2.1-.2.1c-.1.3-.3.7-.6 1c-.3-.3-.5-.7-.6-1c0 0-.1-.1-.2-.1c-.2 0-.2.1-.2.1c-.1.3-.3.7-.6 1c-.3-.3-.5-.7-.6-1c0 0-.1-.1-.2-.1c-.2 0-.2.1-.2.1c-.3.9-1.2 2-2.3 2.1l-.5.7h12.4l-.7-.9"/><path fill="#dbb471" d="M32 23c-5 0-9 4-9 9s4 9 9 9s9-4 9-9s-4-9-9-9m0 16.5c-4.1 0-7.5-3.4-7.5-7.5s3.4-7.5 7.5-7.5s7.5 3.4 7.5 7.5s-3.4 7.5-7.5 7.5"/>`),
		g.Group(children),
	)
}