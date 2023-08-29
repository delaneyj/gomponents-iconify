package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Socialnetwork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M704 1024q-87 0-161-34.5T426.5 896T384 767.5T426.5 639T543 546t161-34q36 0 73 7q21-34 50.5-68t49.5-50l19-17v179q59 36 93.5 89.5T1024 768q0 69-43 128t-116.5 93.5T704 1024zm-32-576q-133 0-214 52T343 639q-44-4-88-18q-21 37-53 74t-53 55l-21 18V559q-60-45-94-107T0 320q0-87 51.5-160.5T191.5 43t193-43t193 43T717 159.5T768 320q0 68-32 129q-30-1-64-1z"/>`),
		g.Group(children),
	)
}