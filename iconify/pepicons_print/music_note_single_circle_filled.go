package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNoteSingleCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M26 14c0 6.627-5.373 12-12 12S2 20.627 2 14S7.373 2 14 2s12 5.373 12 12Z" opacity=".2"/><path fill-rule="evenodd" d="M19.75 10.804c0-1.542-2.146-3.669-6.469-6.606a.5.5 0 0 0-.781.413v3.684a.5.5 0 0 0 .343.475c1.474.487 2.591 1.263 3.365 2.328c.734 1.01.868 1.858.46 2.617c-.236.44.25.917.686.672c1.575-.884 2.396-2.09 2.396-3.583ZM13.5 7.941V5.564c3.492 2.45 5.25 4.259 5.25 5.24c0 .746-.278 1.396-.855 1.96c-.02-.724-.316-1.48-.878-2.254C16.177 9.355 15 8.498 13.5 7.94Z" clip-rule="evenodd"/><path d="M12.5 7.5a.5.5 0 0 1 1 0V18a.5.5 0 0 1-1 0V7.5Z"/><path fill-rule="evenodd" d="M7.5 18c0 1.408 1.365 2.5 3 2.5s3-1.092 3-2.5s-1.365-2.5-3-2.5s-3 1.092-3 2.5Zm5 0c0 .8-.874 1.5-2 1.5s-2-.7-2-1.5s.874-1.5 2-1.5s2 .7 2 1.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}