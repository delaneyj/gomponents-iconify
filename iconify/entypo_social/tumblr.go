package entypo_social

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tumblr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.6 18.196c-.777.371-1.48.631-2.109.781a8.92 8.92 0 0 1-2.043.223c-.831 0-1.566-.107-2.205-.318a4.757 4.757 0 0 1-1.635-.908c-.451-.395-.764-.812-.938-1.254c-.174-.443-.261-1.086-.261-1.926V8.339H4.4V5.735c.714-.234 1.326-.57 1.835-1.01a5.04 5.04 0 0 0 1.227-1.58c.308-.613.519-1.396.636-2.345h2.585v4.652h4.314v2.887h-4.314v4.719c0 1.066.056 1.752.168 2.055c.111.303.319.545.622.725c.403.244.863.367 1.381.367c.92 0 1.836-.303 2.746-.908v2.899z"/>`),
		g.Group(children),
	)
}