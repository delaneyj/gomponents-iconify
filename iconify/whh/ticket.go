package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ticket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024.226 512v192q0 26-19 45t-45 19h-576v-32q0-13-9.5-22.5t-22.5-9.5t-22.5 9.5t-9.5 22.5v32h-256q-26 0-45-19t-19-45V512q53 0 90.5-37.5t37.5-90.5t-37.5-90.5T.226 256V64q0-26 19-45t45-19h256v32q0 13 9.5 22.5t22.5 9.5t22.5-9.5t9.5-22.5V0h576q26 0 45 19t19 45v192q-53 0-90.5 37.5t-37.5 90.5t37.5 90.5t90.5 37.5zm-640-352q0-13-9.5-22.5t-22.5-9.5t-22.5 9.5t-9.5 22.5v64q0 13 9.5 22.5t22.5 9.5t22.5-9.5t9.5-22.5v-64zm0 192q0-13-9.5-22.5t-22.5-9.5t-22.5 9.5t-9.5 22.5v64q0 13 9.5 22.5t22.5 9.5t22.5-9.5t9.5-22.5v-64zm0 192q0-13-9.5-22.5t-22.5-9.5t-22.5 9.5t-9.5 22.5v64q0 13 9.5 22.5t22.5 9.5t22.5-9.5t9.5-22.5v-64z"/>`),
		g.Group(children),
	)
}