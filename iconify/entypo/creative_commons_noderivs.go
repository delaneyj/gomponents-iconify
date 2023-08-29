package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreativeCommonsNoderivs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.988.4c-2.66 0-4.91.928-6.753 2.784C1.345 5.104.4 7.376.4 10c0 2.624.945 4.88 2.835 6.768c1.89 1.888 4.142 2.832 6.753 2.832c2.643 0 4.934-.952 6.872-2.856c1.827-1.808 2.74-4.056 2.74-6.744c0-2.688-.93-4.96-2.788-6.816C14.954 1.328 12.68.4 9.988.4zm.024 1.727c2.179 0 4.029.769 5.55 2.305c1.54 1.52 2.308 3.375 2.308 5.568c0 2.208-.752 4.04-2.258 5.496c-1.586 1.568-3.453 2.351-5.6 2.351c-2.146 0-3.996-.775-5.55-2.328C2.907 13.967 2.13 12.128 2.13 10c0-2.129.785-3.984 2.355-5.568c1.506-1.536 3.348-2.305 5.527-2.305z" clip-rule="evenodd"/><path fill="currentColor" fill-rule="evenodd" d="M13.627 7.725H6.649v1.653h6.978V7.725zm0 3.085H6.649v1.653h6.978V10.81z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}