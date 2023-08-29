package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Asteroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M566 503q-9 9-22.5 9t-22.5-9t-9-22.5t9-22.5L970 9q9-9 22-9t22.5 9.5T1024 32t-9 22zm-192-64q-9 9-22.5 9t-22.5-9t-9-22.5t9-22.5L714 9q9-9 22-9t22.5 9.5T768 32t-9 22zm138 329q0 106-75 181t-181 75t-181-75T0 768t75-181t181-75t181 75t75 181zm458-503q9-9 22-9t22.5 9.5t9.5 22.5t-9 22L630 695q-9 9-22.5 9t-22.5-9t-9-22.5t9-22.5z"/>`),
		g.Group(children),
	)
}