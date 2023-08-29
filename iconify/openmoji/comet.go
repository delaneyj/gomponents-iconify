package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Comet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<ellipse cx="20.767" cy="48.407" fill="#ea5a47" rx="6.586" ry="7.025" transform="rotate(-25.837 20.767 48.407)"/><path fill="#d22f27" d="M24.98 42.16c-.036-.034.104 1.5.114 1.548a9.468 9.468 0 0 1-.56 4.409c-.753 2.321-2.116 5.892-3.164 5.954c-2.231.133 4.166.094 4.166.094s.438-.096.878-.878c.179-.317 1.24-.199 1.419-.813c.172-.592.443-.485.702-.741c.674-.667.544-.578.514-1.08c-.042-.697 0-4.391 0-4.391Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.49" d="m28.054 44.314l1.033 2.742l.312 2.05l-1.016 2.793l-.979 1.171l-1.6.703l-1.082.596l-2.495.693l-.757-.165l-1.964-.025l-2.03-.911l-1.612-2.042l-.749-2.004l-.974-2.012l.448-1.285l.325-1.23l.5-2.333l3.81-2.037L21.56 41s.936.42 1.258.56c.112.049 1.054.38 1.173.408Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m26.841 36.759l19.435-19.436M31.782 37.982l23.6-23.6M35.171 40.229l15.964-15.964m-10.964-6.036L43.7 14.7M25.171 33.229L28.7 29.7m4.471 16.529L36.7 42.7m19.471-23.471L59.7 15.7"/><path d="M22.608 30.496a1.126 1.126 0 0 0 0-2.25a1.126 1.126 0 0 0 0 2.25Zm8 28a1.126 1.126 0 0 0 0-2.25a1.126 1.126 0 0 0 0 2.25Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M40 41h2l.75-2.25l.75 2.25h2L44 42.25l.5 2L42.75 43L41 44.25l.75-2L40 41zM28 20h2l.75-2.25l.75 2.25h2L32 21.25l.5 2L30.75 22L29 23.25l.75-2L28 20z"/><path d="M53.608 28.496a1.126 1.126 0 0 0 0-2.25a1.126 1.126 0 0 0 0 2.25Z"/>`),
		g.Group(children),
	)
}