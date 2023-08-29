package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spinner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M212.5 7Q231 7 244 20.5T257 52t-13 31.5T212.5 97T181 83.5T168 52t13-31.5T212.5 7zM337 54q22 0 37.5 16t15.5 37.5t-15.5 37.5t-37.5 16t-38-16t-16-37.5T299 70t38-16zM91.5 64q16.5 0 28 12t11.5 28.5t-11.5 28.5t-28 12T63 133t-12-28.5T63 76t28.5-12zM34 198q14 0 24 10t10 24t-10 24t-24 10t-24-10t-10-24t10-24t24-10zm355 0q14 0 24 10t10 24t-10 24t-24 10t-24-10t-10-24t10-24t24-10zM85 325q14 0 24 10t10 24t-10 24t-24 10t-24-10t-10-24t10-24t24-10zm255 0q14 0 24 10t10 24t-10 24t-24 10t-24-10t-10-24t10-24t24-10zm-128 47q14 0 24.5 10.5T247 407t-10.5 24t-24.5 10t-24-10t-10-24t10-24.5t24-10.5z"/>`),
		g.Group(children),
	)
}