/*
 * @Author: lwnmengjing
 * @Date: 2021/6/24 10:49 上午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/24 10:49 上午
 */

package form

type Pagination struct {
	Page     int `form:"page" query:"page"`
	PageSize int `form:"pageSize" query:"pageSize"`
}
