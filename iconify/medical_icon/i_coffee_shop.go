package medical_icon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ICoffeeShop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M50.409 14.564c-.171 0-.341.004-.51.011v-.009H14.628v23.245a6.296 6.296 0 0 0 6.295 6.296h22.679a6.298 6.298 0 0 0 6.296-6.296v-1.656c.168.007.338.013.509.013c6.199 0 11.224-4.665 11.224-10.862c.001-6.198-5.024-10.742-11.222-10.742zm0 17.326c-.171 0-.341-.009-.51-.021V18.814c.167-.01.339-.018.51-.018c3.756 0 6.801 2.754 6.801 6.511s-3.044 6.583-6.801 6.583zm12.462 13.575v.218a6.297 6.297 0 0 1-6.295 6.296H7.59a6.296 6.296 0 0 1-6.295-6.296v-.218h61.577z"/>`),
		g.Group(children),
	)
}