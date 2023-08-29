package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonPlusCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><g opacity=".2"><path fill-rule="evenodd" d="M15.95 10.75a1 1 0 0 1 1-1h5.25a1 1 0 1 1 0 2h-5.25a1 1 0 0 1-1-1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M19.5 14.5a1 1 0 0 1-1-1V8.25a1 1 0 1 1 2 0v5.25a1 1 0 0 1-1 1Z" clip-rule="evenodd"/><path d="M14 9.6c0 1.436-1.12 2.6-2.5 2.6S9 11.036 9 9.6C9 8.164 10.12 7 11.5 7S14 8.164 14 9.6Z"/><path fill-rule="evenodd" d="M11.5 8c-.792 0-1.5.679-1.5 1.6s.708 1.6 1.5 1.6s1.5-.679 1.5-1.6S12.292 8 11.5 8ZM8 9.6C8 7.65 9.53 6 11.5 6S15 7.65 15 9.6c0 1.95-1.53 3.6-3.5 3.6S8 11.55 8 9.6Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M16.5 20v-2.167c0-2.684-2.254-4.766-4.987-4.766c-2.732 0-5.013 2.082-5.013 4.766L6.502 20" clip-rule="evenodd"/><path fill-rule="evenodd" d="M5.5 17.833c0-3.295 2.79-5.766 6.013-5.766c3.232 0 5.987 2.478 5.987 5.766V20a1 1 0 1 1-2 0v-2.167c0-2.08-1.753-3.766-3.987-3.766c-2.24 0-4.013 1.692-4.013 3.766l.002 2.166a1 1 0 0 1-2 .002L5.5 17.833Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M16.5 21h-10v-1h10v1Z" clip-rule="evenodd"/></g><path fill-rule="evenodd" d="M14.95 9.75a.5.5 0 0 1 .5-.5h5.25a.5.5 0 1 1 0 1h-5.25a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M18 13a.5.5 0 0 1-.5-.5V7.25a.5.5 0 0 1 1 0v5.25a.5.5 0 0 1-.5.5ZM9.8 6a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5ZM6.3 8.5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M3.8 17.5c0-3.322 2.67-6.5 6-6.5s6 3.178 6 6.5v2a.5.5 0 0 1-1 0v-2c0-2.873-2.32-5.5-5-5.5s-5 2.627-5 5.5v2a.5.5 0 0 1-1 0v-2Z" clip-rule="evenodd"/><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}