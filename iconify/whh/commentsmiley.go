package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Commentsmiley(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 896q-66 0-134-16q-34 40-69.5 69.5t-60 43.5t-47.5 21.5t-30.5 8.5t-10.5 1q26-57 30-124.5T177 786Q94 723 47 635T0 448q0-91 40.5-174t109-143T313 35.5T512 0t199 35.5T874.5 131t109 143t40.5 174t-40.5 174t-109 143T711 860.5T512 896zm-64-608q0-13-9.5-22.5T416 256h-64q-13 0-22.5 9.5T320 288v64q0 13 9.5 22.5T352 384h64q13 0 22.5-9.5T448 352v-64zm0 256q0-13-9.5-22.5T416 512h-64q-13 0-22.5 9.5T320 544v64q0 13 9.5 22.5T352 640h64q13 0 22.5-9.5T448 608v-64zm192-352h-64q23 23 37 49t19.5 63.5T639 368t1 80t-1 80t-6.5 63.5T613 655t-37 49h64q26-26 39-39.5t33-40.5t30-49.5t18-56t8-70.5t-8-70.5t-18-56t-30-49t-33-40.5t-39-40z"/>`),
		g.Group(children),
	)
}