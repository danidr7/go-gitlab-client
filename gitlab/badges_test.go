package gitlab

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGitlab_ProjectBadges(t *testing.T) {
	ts, gitlab := mockServerFromMapping(t, "badges/project_1_badges.json")
	defer ts.Close()

	badges, meta, err := gitlab.ProjectBadges("1", nil)

	assert.NoError(t, err)

	assert.Equal(t, 5, len(badges.Items))

	assert.IsType(t, new(ResponseMeta), meta)
	assert.Equal(t, 1, meta.Page)
	assert.Equal(t, 10, meta.PerPage)
}

func TestGitlab_ProjectBadge(t *testing.T) {
	ts, gitlab := mockServerFromMapping(t, "badges/project_1_badge_1.json")
	defer ts.Close()

	badge, meta, err := gitlab.ProjectBadge("1", 1)

	assert.NoError(t, err)

	assert.IsType(t, new(Badge), badge)
	assert.Equal(t, 1, badge.Id)

	assert.IsType(t, new(ResponseMeta), meta)
}