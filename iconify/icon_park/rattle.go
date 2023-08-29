package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rattle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="30.075" cy="15.562" r="11" fill="#2F88FF" transform="rotate(40 30.075 15.562)"/><path d="M21.6485 8.49117C17.0389 8.53954 15.322 5.1407 13.0722 7.82186C11.0395 10.2443 13.5652 13.4572 11.6369 15.7553"/><path d="M25.3022 25.9167C25.3022 25.9167 21.1241 30.8959 19.5171 32.8111C17.9101 34.7262 17.8352 37.9269 15.9069 40.225C13.9785 42.5231 11.1609 42.7696 8.86272 40.8413C6.56459 38.9129 6.31808 36.0952 8.24644 33.7971C10.1748 31.499 13.3139 30.8694 14.9208 28.9543C16.5278 27.0392 20.0837 23.0608 20.7059 22.0599"/><circle cx="11.24" cy="19.339" r="3" fill="#2F88FF" transform="rotate(40 11.24 19.34)"/><circle cx="28.462" cy="37.707" r="3" fill="#2F88FF" transform="rotate(40 28.462 37.707)"/><path d="M37.2158 24.165C37.8454 27.304 40.0686 32.4331 38.0786 34.0268C36.0886 35.6205 31.554 32.4682 30.3917 35.4091"/></g>`),
		g.Group(children),
	)
}