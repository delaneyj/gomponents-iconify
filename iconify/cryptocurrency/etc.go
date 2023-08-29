package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Etc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm6.732-16L16 12.19L9.268 16L16 19.781l6.732-3.783zM16 21.047a3944.37 3944.37 0 0 0-7-3.952c2.079 3.248 4.66 7.26 7 10.904c2.34-3.643 4.921-7.656 7-10.904a3944.185 3944.185 0 0 0-7 3.952zm0-10.089l7 3.907L16 4L9 14.866l7-3.907z"/><path fill-opacity=".296" fill-rule="nonzero" d="m22.71 15.976l-6.721.577v-4.379l6.72 3.802zm-6.721 5.038c1.98-1.12 4.537-2.564 6.988-3.944c-2.076 3.242-4.652 7.246-6.988 10.882v-6.938zm0-10.069V4l6.988 10.845l-6.988-3.9z"/><path fill-opacity=".803" d="m15.989 16.553l6.72-.577l-6.72 3.775z"/><path d="m15.988 16.553l-6.721-.577l6.721 3.775z" opacity=".311"/></g>`),
		g.Group(children),
	)
}