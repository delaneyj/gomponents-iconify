package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Voicemail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M394.5 64q48.5 0 83 34.5t34.5 83t-34.5 83T395 299H117q-48 0-82.5-34.5T0 181.5t34.5-83t83-34.5t83 34.5T235 181q0 43-27 75h96q-27-32-27-75q0-48 34.5-82.5t83-34.5zM117 256q31 0 53-22t22-53t-22-52.5t-53-21.5t-52.5 21.5T43 181t21.5 53t52.5 22zm278 0q31 0 52.5-22t21.5-53t-21.5-52.5T395 107t-53 21.5t-22 52.5t22 53t53 22z"/>`),
		g.Group(children),
	)
}