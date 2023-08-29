package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PopTime(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.837 5.712a12.627 12.627 0 0 0-5.232-1.211a11.567 11.567 0 0 0-.157 2.41c-2.692-.845-7.701-.25-7.701-.25a43.662 43.662 0 0 1 .751 6.7c.313 7.17-3.475 9.361-3.475 15.216s3.726 8.797 8.923 8.797c1.659 0 3.049-1.66 2.548-2.88m7.671-26.332c1.775 1.955 3.118 4.954 3.118 9.55c0 5.229 3.005 4.415 3.005 11.303s-5.76 8.359-8.578 8.359c-2.505 0-2.567-2.818-2.567-2.818"/><circle cx="23.656" cy="28.091" r="7.905" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.75 12.872a4.083 4.083 0 1 0-1.257-4.253"/><circle cx="19.834" cy="11.716" r="4.084" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.235 14.097a14.747 14.747 0 0 1 0 3.193"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.863 16.789c-.532 1.076-2.646 1.588-4.743.68"/><circle cx="20.822" cy="13.173" r=".75" fill="currentColor"/><circle cx="25.949" cy="11.716" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.831 32.076c2.197.132 4.277 1.353 4.277 4.077s-3.739 4.525-3.739 4.525c3.6.751.874 2.816-.3 1.642c.063 1.347-2.223 1.785-2.943 0c-.5 1.472-2.285 1.785-2.254-.407c-1.096.376-2.504-.281-1.534-1.659c0 0-3.788.376-5.541.376s-5.135-.376-5.135-.376c.97 1.378-.438 2.035-1.534 1.66c.031 2.19-1.753 1.878-2.254.406c-.72 1.785-3.006 1.347-2.943 0c-1.174 1.174-3.9-.89-.3-1.642c0 0-3.738-1.802-3.738-4.525s1.47-3.945 3.667-4.077"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.587 32.584a4.083 4.083 0 0 1 1.765 1.377a13.07 13.07 0 0 1 .1-1.83m-11.772.453a4.083 4.083 0 0 0-1.765 1.377a13.049 13.049 0 0 0-.1-1.83"/>`),
		g.Group(children),
	)
}