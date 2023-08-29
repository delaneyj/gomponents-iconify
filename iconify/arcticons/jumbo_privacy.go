package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JumboPrivacy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.038 40.042a4.253 4.253 0 0 1-.856.908c-2.754 2.09-5.871 2.2-8.223 2.2a2.048 2.048 0 0 1-2.346-2.098a1.677 1.677 0 0 1 1.83-1.88c1.357 0 3.193-.287 3.193-2.372s-2.084-2.123-3.442-2.123h-5.848M8.331 9.278a27.234 27.234 0 0 1 14.283-3.823c-.612 3.213-1.53 3.366-1.53 3.366s3.111-1.938 8.517-1.224a5.095 5.095 0 0 1-2.397 3.264s7.038.408 7.7 9.843a58.942 58.942 0 0 1 .1 6.52"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.543 12.823c3.464-.674 11.42-.942 11.42 8.123a15.259 15.259 0 0 1-5.392 12.185"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.68 28.641c2.095-1.384 3.304-3.632 3.304-7.121c0-3.825-3.884-4.686-4.83-4.676M20.919 31.041c.115 1.014 1.568 3.117 3.385.86m-9.316-10.644c.402-.574 1.855-1.664 3.27.363m11.23-.008c-.44-.707-1.664-1.797-3.27.306"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.11 39.509c4.607-1.973 10.937-7.257 11.42-11.028c.593-4.628-4.561-7.363-6.646-7.166c-2.416.227-6.699 7.407-9.759 10.9m30.298 10.594s-.087.048-1.245-1.36C31.412 39.3 24.162 29.6 23.583 27.793a3.538 3.538 0 0 1 1.734-4.106c1.122-.637 4.54-3.009 6.4-1.657c1.718 1.247 5.518 9.949 7.278 15.385c.547 1.688.491 1.498.491 1.498"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.903 22.455c-.017.652 1.532 2.354 3.788.71c-.038 1.606 1.32 2.715 3.29 2.237c-.192 1.683-.303 3.066 1.388 3.847m3.12-1.819c1.19.247 3.371-.288 2.568-3.023a2.035 2.035 0 0 0 3.347-1.32c1.147 1.32 2.216 1.593 3.623.548"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}