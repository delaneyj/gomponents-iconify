package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Subrion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 384q-87 0-160.5 43T491 543.5T448 704q0 85 43 160q-69 18-130 47.5t-94.5 53t-70 41.5t-68.5 18q-56 0-92-50T0 832q0-61 33-134t92-141q-28-17-44.5-46T64 448q0-81 92-151.5T392 202q-8-23-8-42q0-66 56-113T576 0q90 0 171 41.5t139.5 111t94.5 163t42 195.5q-45-59-111.5-93T768 384zm0 64q106 0 181 75t75 181t-75 181t-181 75t-181-75t-75-181t75-181t181-75zm64.5 256q26.5 0 45-18.5t18.5-45t-18.5-45.5t-45-19t-45.5 19t-19 45.5t19 45t45.5 18.5z"/>`),
		g.Group(children),
	)
}