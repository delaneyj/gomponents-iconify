package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Disc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><g transform="translate(3 3)"><path d="M7.466 1.066L8.984-.452l.869.87l-1.518 1.517zM.895 10.405l-.868-.87L1.545 8.02l.868.87z"/><ellipse cx="5.115" cy="5.141" rx="3.115" ry="3.141"/><path fill="currentColor" d="M4.988 3.022a1.977 1.977 0 1 0 0 3.953a1.978 1.978 0 0 0 1.979-1.976a1.98 1.98 0 0 0-1.979-1.977z"/></g><path fill="currentColor" d="M15.958 7.958C15.958 3.516 12.385 0 7.979 0C3.573 0 0 3.516 0 7.958S3.572 16 7.979 16c4.564 0 7.979-3.69 7.979-8.042zm-3.697-5.067l.869.869l-1.518 1.518l-.869-.869l1.518-1.518zM3.633 13.193l-.868-.869l1.518-1.518l.868.869l-1.518 1.518zm1.252-5.208C4.885 6.25 6.28 4.843 8 4.843c1.719 0 3.115 1.406 3.115 3.142c0 1.733-1.396 3.14-3.115 3.14c-1.721 0-3.115-1.406-3.115-3.14z"/></g>`),
		g.Group(children),
	)
}