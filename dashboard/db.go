package dashboard

import (
	"fmt"

	"gorm.io/gorm"

	"devex_dashboard/project"
)

func gitChangesData(db *gorm.DB, filesMode bool, projects []project.ID, filesFilter string) (barNames []string, result []timedData, err error) {
	var bars values

	// Future: months/weeks selector

	grouping := "alias, package"
	barFilter := "alias || '/' || package"

	if filesMode {
		grouping += ", name"
		filesFilter = "and f.present > 0\n" + filesFilter
		barFilter += " || '/' || name"
	}

	sqlBars := `
	with fcm as (select %[1]s, date("time", 'start of month') as month, sum(rows_added + rows_removed) as line_changes
		from git_changes as ch
		join files f on ch.file = f.id
		join projects p on f.project = p.id
		where f.project in ?
		   %[2]s
		   and time > date('now', '-48 month')
		group by %[1]s, date("time", 'start of month'))
	select %[1]s, count(*), sum(line_changes), avg(line_changes) as value
	from fcm group by %[1]s
	having count(*) > 6
	order by avg(line_changes) desc
	limit 20
`
	sqlBars = fmt.Sprintf(sqlBars, grouping, filesFilter)

	err = db.Raw(sqlBars, projects).Scan(&bars).Error
	if err != nil {
		return nil, nil, err
	}

	barNames = bars.barNames()

	sql := `
	select %[1]s, date("time", 'start of month') as bar_time, sum(rows_added + rows_removed) as value
	from git_changes as ch
	join files f on ch.file = f.id
	join projects p on f.project = p.id
	where f.project in ?
		and %[2]s in ?
		and time > date('now', '-24 month')
	group by %[1]s, date("time", 'start of month')
`
	sql = fmt.Sprintf(sql, grouping, barFilter)

	err = db.Raw(sql, projects, barNames).Scan(&result).Error

	return barNames, result, err
}
