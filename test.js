function ptr(arrList, left, right) {
    let first = arrList[left]
    while (left < right) {
        while (left < right && arrList[right] > first) {
            right -= 1
        }
        arrList[left] = arrList[right]
        while (left < right && arrList[left] < first){
            left += 1
        }
        arrList[right] = arrList[left]
    }
    first = arrList[left]
    return first
}


function sort(arrList, left, right) {
    mid = ptr(arrList, left, right)
    sort(arrList, 0, mid-1)
    sort(arrList, mid+1 ,arrList.length())
}
