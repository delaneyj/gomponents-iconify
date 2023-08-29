package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonPlusOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g opacity=".2"><path fill-rule="evenodd" d="M12.95 7.75a1 1 0 0 1 1-1h5.25a1 1 0 1 1 0 2h-5.25a1 1 0 0 1-1-1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M16.5 11.5a1 1 0 0 1-1-1V5.25a1 1 0 1 1 2 0v5.25a1 1 0 0 1-1 1Z" clip-rule="evenodd"/><path d="M11 6.6c0 1.436-1.12 2.6-2.5 2.6S6 8.036 6 6.6C6 5.164 7.12 4 8.5 4S11 5.164 11 6.6Z"/><path fill-rule="evenodd" d="M8.5 5C7.708 5 7 5.679 7 6.6s.708 1.6 1.5 1.6S10 7.521 10 6.6S9.292 5 8.5 5ZM5 6.6C5 4.65 6.53 3 8.5 3S12 4.65 12 6.6c0 1.95-1.53 3.6-3.5 3.6S5 8.55 5 6.6Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13.5 17v-2.167c0-2.684-2.254-4.766-4.987-4.766c-2.732 0-5.013 2.082-5.013 4.766L3.502 17" clip-rule="evenodd"/><path fill-rule="evenodd" d="M2.5 14.833c0-3.295 2.79-5.766 6.013-5.766c3.232 0 5.987 2.478 5.987 5.766V17a1 1 0 1 1-2 0v-2.167c0-2.08-1.753-3.766-3.987-3.766c-2.24 0-4.013 1.692-4.013 3.766l.002 2.166a1 1 0 0 1-2 .002L2.5 14.833Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13.5 18h-10v-1h10v1Z" clip-rule="evenodd"/></g><path fill-rule="evenodd" d="M11.95 6.75a.5.5 0 0 1 .5-.5h5.25a.5.5 0 1 1 0 1h-5.25a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M15 10a.5.5 0 0 1-.5-.5V4.25a.5.5 0 0 1 1 0V9.5a.5.5 0 0 1-.5.5ZM6.8 3a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5ZM3.3 5.5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M.8 14.5c0-3.322 2.67-6.5 6-6.5s6 3.178 6 6.5v2a.5.5 0 0 1-1 0v-2c0-2.873-2.32-5.5-5-5.5s-5 2.627-5 5.5v2a.5.5 0 0 1-1 0v-2Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}