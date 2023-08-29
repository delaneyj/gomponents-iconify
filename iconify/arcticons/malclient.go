package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Malclient(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.944 20.212c1.617-.594 3.313.464 4.068 2.536a6.254 6.254 0 0 1-.96 6.057m-1.543.905c-1.477.397-2.974-.544-3.733-2.347a6.487 6.487 0 0 1 .361-5.665M43.5 18.87a121.763 121.763 0 0 0-39-5.978c14.304.023 24.25 3.04 35.312 11.007"/><ellipse cx="24.95" cy="25.687" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="6.627" ry="9.42"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.844 35.074a32.243 32.243 0 0 0 10.191-3.129m-8.962-3.129a1.244 1.244 0 0 1 .484 2.262a1.26 1.26 0 0 1-1.397-.008a1.235 1.235 0 0 1-.51-1.296m-1.818-9.562a1.254 1.254 0 1 1-1.54 1.41m-6.465-2.83c-1.2.369-9.404 5.302-9.404 5.302l9.943 6.407l-4.17-6.387l3.63-5.322"/>`),
		g.Group(children),
	)
}