
1. git 找出 那些文件是谁创建的
git ls-tree --name-only HEAD | while read filename; do     echo "$(git log --pretty=format:"%an" -- $filename | tail -1) $filename"; done | grep Guohong

