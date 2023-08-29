package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloseUpMode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M12.5 0C11.576.997 8.951 2.519 9 7c-.448-1.025-.976-3.303-.813-5.031c0 0-4.187 1.788-4.187 6.437c0 3.483 3.02 5.993 7 6.5V26h3c4.392-.11 12.855-2.495 11.813-13.969c-4.799.71-10.605 5.749-11.813 8.938v-6.063c3.98-.506 7-3.016 7-6.5c0-4.649-4.188-6.437-4.188-6.437C16.977 3.697 16.448 5.975 16 7c.049-4.481-2.576-6.003-3.5-7zM.062 12.094s-.198 4.148.844 7.5c1.218 3.913 4.722 6.375 9.094 6.375v-5.407C6.78 14.989.062 12.095.062 12.095z"/>`),
		g.Group(children),
	)
}