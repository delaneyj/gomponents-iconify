package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Album(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.193 1024h-896q-27 0-45.5-19t-18.5-45V64q0-26 18.5-45t45.5-19h896q26 0 45 19t19 45v896q0 27-19 45.5t-45 18.5zm0-896q0-26-19-45t-45-19h-704q-27 0-45.5 18.5t-18.5 45.5v768q0 27 18.5 45.5t45.5 18.5h704q26 0 45-18.5t19-45.5V128zm-416 768q-96 0-177-47t-128-128t-47-177t47-177t128-128t177-47t177 47t128 128t47 177t-47 177t-128 128t-177 47zm0-512q-66 0-113 47t-47 113t47 113t113 47t113-47t47-113t-47-113t-113-47zm0 256q-40 0-68-28t-28-68t28-68t68-28t68 28t28 68t-28 68t-68 28z"/>`),
		g.Group(children),
	)
}