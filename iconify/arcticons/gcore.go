package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gcore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" d="M5.5 19.7a73.46 73.46 0 0 1 13.434-.6"/><path fill="none" stroke="currentColor" stroke-linecap="round" d="M12.999 6.641c.068 19.781 0 12.848.408 34.056m-.245-10.437a19.343 19.343 0 0 1-3.41 5.247c-.215.25-.407.417-.623.667c-.203.235-.893.822-.893.822m5.123-3.706a79.083 79.083 0 0 1 3.869 2.235M27.07 5.758a36.576 36.576 0 0 0 3.33 6.933m-13.662 2.447c6.088-.64 21.668-1.225 25.762-.68M29.075 16.57c-8.43 10.14-13.902 14.537.384 16.63"/><path fill="none" stroke="currentColor" stroke-linecap="round" d="M34.628 26.573s-2.788 4.999-9.806 11.536m9.71-3.173a36.959 36.959 0 0 1-6.345 6.826m3.268-2.692a6.737 6.737 0 0 0 4.038 3.172"/>`),
		g.Group(children),
	)
}