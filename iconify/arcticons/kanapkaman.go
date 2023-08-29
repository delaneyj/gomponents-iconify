package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kanapkaman(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.94 15.79c5.418-1.108 18.552-3.78 19.217 4.366a63.315 63.315 0 0 1-15.704 5.217C18.388 26.73 7.94 28.43 6.65 23.949c-1.384-4.805 8.091-6.482 16.29-8.158Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.708 14.575c1.056.285 1.518 2.147 1.6 3.558M22.94 15.79c.851.4 1.73 2.955 1.756 3.834m-9.708-2.098a5.014 5.014 0 0 1 2.149 4.067m25.02-1.437c.346.985 1.53 1.61 1.318 2.702s-2.276 1.038-2.436 1.663s-.413 1.544-.945 1.823s-1.75-.286-2.222 0s-1.222 1.792-2.183 2.183c-1.111.452-1.936-2.083-2.861-1.877s-1.957 2.722-3.008 2.702s-1.97-2.163-3.234-2.236s-1.657 2.168-3.18 2.236c-1.046.047-1.378-1.382-2.037-1.357c-2.053.076-3.97 3.005-5.736 3.06c-2.865.09-3.444-2.425-4.884-2.661c-.947-.155-2.033.935-3.114.652c-.597-.156-.163-1.742-1.12-2.195s-1.622.538-2.007-.833c-.126-.446 1.457-1.75 2.142-2.07"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.655 26.74c.017 1.5-.35 5.048 1.993 6.112c3.282 1.492 13.435.12 18.645-.399s14.673-3.114 15.451-4.95a4.528 4.528 0 0 0 0-3.466"/>`),
		g.Group(children),
	)
}