package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zodiacleo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M864 1024q-93 0-158.5-84.5T640 736q0-11 1-28.5t9-60.5t22-71q32-64 32-160v-64q0-93-65.5-158.5T480 128t-158.5 65.5T256 352q0 52 22.5 97.5T341 527q76 27 123.5 92.5T512 768q0 106-75 181t-181 75t-181-75T0 768q0-83 48.5-149.5T174 526q-46-81-46-174q0-96 47-177T303 47T480 0t177 47t128 128t47 177v96q0 33-32 114.5T768 672v64q0 66 28 113t68 47q43 0 96-44v150q-51 22-96 22zM128 768q0 53 37.5 90.5T256 896t90.5-37.5T384 768t-37.5-90.5T256 640t-90.5 37.5T128 768z"/>`),
		g.Group(children),
	)
}