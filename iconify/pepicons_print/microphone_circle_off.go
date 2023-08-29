package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrophoneCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M14.5 5.5A2.5 2.5 0 0 0 12 8v4a2.5 2.5 0 0 0 5 0V8a2.5 2.5 0 0 0-2.5-2.5ZM10 21c0 .792 1.666 1 4.5 1s4.5-.208 4.5-1c0-.69-1.264-.937-3.464-.989v-1.588C18.333 18 20.5 15.9 20.5 13.3v-1.8a1 1 0 1 0-2 0v1.8c0 1.725-1.756 3.2-4 3.2c-2.244 0-4-1.475-4-3.2v-1.8a1 1 0 0 0-2 0v1.8c0 2.623 2.204 4.737 5.036 5.133v1.576C11.292 20.058 10 20.303 10 21Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M10.75 7.25a2.25 2.25 0 0 1 4.5 0v3.5a2.25 2.25 0 0 1-4.5 0v-3.5Z" clip-rule="evenodd"/><path d="M13 20c-2.48 0-4-.217-4-1s1.52-1 4-1s4 .217 4 1s-1.52 1-4 1Z"/><path d="M12.5 15.5h1V19h-1v-3.5Z"/><path d="M17 10.5a.5.5 0 0 1 1 0v1.65c0 2.421-2.254 4.35-5 4.35s-5-1.929-5-4.35V10.5a.5.5 0 0 1 1 0v1.65c0 1.831 1.775 3.35 4 3.35s4-1.519 4-3.35V10.5Z"/><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}