package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Supertux(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.887 15.047c.631-5.263-2.863-8.926-8.961-8.926c-6.05 0-8.829 4.042-8.829 8.969S6.634 23.3 6.634 30.079c0 2.947 3.937 8.442 7.632 8.442c5.478 0 7.12 1.69 12.82.813"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.236 26.004a8.052 8.052 0 0 1 7.04-5.567c2.779 0 7.668 3.468 7.668 10.584a8.428 8.428 0 0 1-8.51 8.321c-2.512 0-5.992-1.269-6.724-6.838"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.581 22.858a13.571 13.571 0 0 0 5.813 4.208c1.83.781 3.241 1.88 3.03 3.013c-.421 2.253-3.895 2.442-7.095 2.442c-4.653 0-6.632-6.716-6.632-6.716m20.357-2.851c-1.167-1.517-2.557-2.717-2.652-4.296"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.992 16.574s1.295-1.927 3.158-1.927s7.674 1.264 7.674 2.495s-6.38 1.8-7.958 1.8c-1.453 0-2.874-2.368-2.874-2.368Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.792 16.921c2.052.916 8.715.678 8.715.678m.97 9.809c.5-1.092 1.257-2.482 2.378-2.055s1.8 2.337 1.8 2.337s1.595-.49 2.258.71s-.079 1.626-.079 1.626s1.037-.17 1.468.269c.264.268.355.67-.315 1.752c-2.053 3.316-5.527 4.09-7.99 6.71c-1.63 1.735-4.468 5.07-6.853.585m-21.51-9.263c-1.263.316-2.842 1.2-1.79 3.01s4.337 8.497 6.947 8.497c1.537 0 2.337-.117 2.337-1.107a3.082 3.082 0 0 0-.99-2.119"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.718 36.408A121.52 121.52 0 0 1 5.21 38.6a.496.496 0 0 0 .242.756a7.688 7.688 0 0 0 3.444.23m4.796 1.744c1.153.855 4.037.686 4.037-.893a1.696 1.696 0 0 0-1.247-1.817"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.897 39.132c.484 1.01.569 2.532-.842 2.708a3.094 3.094 0 0 1-2.08-.3m2.595-28.724c0 1.79.626 3.179 2.748 3.179a2.87 2.87 0 0 0 3.074-2.926a2.41 2.41 0 0 0-2.632-2.674c-2.147 0-3.19 1.116-3.19 2.42Zm11.35.188c-.033-.485-.552-2.21-1.97-2.21c-.968 0-1.39.696-1.39 2s.592 1.912 1.11 1.996m.57-1.339v-.939m-6.415 1.009v-.939"/>`),
		g.Group(children),
	)
}