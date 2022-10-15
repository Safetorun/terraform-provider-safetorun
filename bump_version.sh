CURRENT_TAG=`git describe --tags --abbrev=0` # e.g. v0.1.18
echo "Current tag $CURRENT_TAG"

IFS='.' # Set delim
read -ra TAG_ARRAY <<< "$CURRENT_TAG" # Read to array
IFS=' '
NEW_VERSION=$((TAG_ARRAY[2]+=1)) # Increment version
NEW_TAG="${TAG_ARRAY[0]}.${TAG_ARRAY[1]}.$NEW_VERSION" # e.g. v0.1.18

echo "New tag $NEW_TAG"

git config user.email "pipe@safetorun.com"
git config user.name "GH Actions"
git remote set-url origin git@github.com:safetorun/safe_to_run_admin_api
git tag $NEW_TAG
git push --tags
