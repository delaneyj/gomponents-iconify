package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreativeCommonsAttribution(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.583 7.623a.62.62 0 0 0-.62-.62H8.037a.62.62 0 0 0-.62.62v3.927h1.095v4.65h2.976v-4.65h1.095V7.623z"/><circle cx="10" cy="5.143" r="1.343" fill="currentColor"/><path fill="currentColor" fill-rule="evenodd" d="M9.988.4c-2.66 0-4.91.928-6.753 2.784C1.345 5.104.4 7.376.4 10c0 2.624.945 4.88 2.835 6.768c1.89 1.888 4.142 2.832 6.753 2.832c2.643 0 4.934-.952 6.872-2.856c1.827-1.808 2.74-4.056 2.74-6.744s-.93-4.96-2.788-6.816C14.954 1.328 12.68.4 9.988.4zm.024 1.728c2.179 0 4.029.768 5.55 2.304c1.54 1.52 2.308 3.376 2.308 5.568c0 2.208-.752 4.04-2.258 5.496c-1.586 1.568-3.453 2.352-5.6 2.352c-2.146 0-3.996-.776-5.55-2.328C2.907 13.968 2.13 12.128 2.13 10c0-2.128.785-3.984 2.355-5.568c1.506-1.536 3.348-2.304 5.527-2.304z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}