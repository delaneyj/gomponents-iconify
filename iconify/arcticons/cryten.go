package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cryten(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" d="M34.946 42.615a54.699 54.699 0 0 1-5.202.699a31.641 31.641 0 0 1-6.663.016c-2.624-.317-5.123-.96-7.118-2.853a9.063 9.063 0 0 1-2.73-5.861c-.365-3.544-.042-7.102-.144-9.714c0-3.995-.023-7.052.01-10.108a11.686 11.686 0 0 1 .846-4.295a8.989 8.989 0 0 1 6.619-5.387a25.205 25.205 0 0 1 8.182-.496c1.984.15 5.914.71 5.914.71"/>`),
		g.Group(children),
	)
}