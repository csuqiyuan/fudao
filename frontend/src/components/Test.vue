<template>
    <div class="plane">
        <div>
            <el-select v-model="grade" placeholder="请选择" @change="selectOneChange()">
                <el-option
                        v-for="(item, index) in grades"
                        :key="index"
                        :label="gradesMap[item.grade]"
                        :value="item.grade">
                </el-option>
            </el-select>
            <el-select v-model="subject" placeholder="请选择" @change="getSubject()">
                <el-option
                        v-for="(item, index) in subjects"
                        :key="index"
                        :label="subjectsMap[item]"
                        :value="item">
                </el-option>
            </el-select>
        </div>

        <div class="table">
            <el-table style="width: 100%" border :data="tableData">
                <template v-for="(item,index) in tableHead">
                    <el-table-column :prop="item.column_name" :label="item.column_comment"
                                     :key="index"></el-table-column>
                </template>
            </el-table>
        </div>

    </div>
</template>

<script>
    export default {
        name: 'Test',
        data() {
            return {
                gradesMap: {
                    "8001": "小班",
                    "8002": "中班",
                    "8003": "大班",
                    "7001": "一年级",
                    "7002": "二年级",
                    "7003": "三年级",
                    "7004": "四年级",
                    "7005": "五年级",
                    "7006": "六年级",
                    "6001": "初一",
                    "6002": "初二",
                    "6003": "初三",
                    "5001": "高一",
                    "5002": "高二",
                    "5003": "高三"
                },
                subjectsMap: {
                    "6001": "语文",
                    "6002": "数学",
                    "6003": "化学",
                    "6004": "物理",
                    "6005": "英语",
                    "6006": "生物",
                    "6007": "政治",
                    "6008": "历史",
                    "6009": "地理",
                    "6010": "讲座",
                    "7057": "编程",
                    "7058": "科学"
                },
                tableHead1: [
                    {
                        column_name: "title", column_comment: "课程名"
                    },
                    {
                        column_name: "course_bgtime", column_comment: "开课时间"
                    },
                    {
                        column_name: "course_endtime", column_comment: "结课时间"
                    },
                    {
                        column_name: "sold_count", column_comment: "购买人数"
                    },
                    {
                        column_name: "course_min_price", column_comment: "价格(元)"
                    },
                    {
                        column_name: "SumPrice", column_comment: "总收入(元)"
                    },
                ],
                tableHead2: [
                    {
                        column_name: "name", column_comment: "课程名"
                    },
                    {
                        column_name: "time_plan", column_comment: "开课计划"
                    },
                    {
                        column_name: "apply_num", column_comment: "预约人数"
                    }
                ],
                tableHead: [],
                tableData: [],
                grades: [],
                subjects: [],
                grade: '',
                subject: ''
            }
        },
        mounted() {
            this.getData();
            console.log(this.formatDate(1230999938))
        },
        methods: {
            getData() {
                var that = this;
                this.$ajax.get('/api/get_grades')
                    .then(function (res) {
                        that.grades = res.data.result["grade_subjects"];
                        console.log(that.grades)
                    })
                    .catch(function (res) {
                        console.log(res)
                    })
            },
            getSubject() {
                var that = this;
                this.$ajax.get('/api/get_subject?grade=' + that.grade + '&subject=' + that.subject)
                    .then(function (res) {
                        if (that.subject === 6010) {
                            console.log(res.data["spe_course_list"].data);
                            that.tableData = res.data["spe_course_list"].data;
                            that.tableHead = that.tableHead2;
                        } else {
                            that.tableData = res.data;
                            that.tableHead = that.tableHead1;
                            for (var i = 0; i < that.tableData.length; i++) {
                                that.tableData[i]["course_min_price"] = that.tableData[i]["course_min_price"] / 100;
                                that.tableData[i]["SumPrice"] = that.tableData[i]["SumPrice"] / 100;

                                var bgtime = that.tableData[i]["course_bgtime"];
                                var bgTimestamp = new Date( bgtime*1000 ) ;
                                that.tableData[i]["course_bgtime"] = bgTimestamp.toLocaleString();

                                var endtime = that.tableData[i]["course_endtime"];
                                var endTimestamp = new Date( endtime*1000 ) ;
                                that.tableData[i]["course_endtime"] = endTimestamp.toLocaleString();
                            }
                        }

                    })
                    .catch(function (res) {
                        console.log(res)
                    });
            },
            selectOneChange() {
                var that = this;
                for (var i = 0; i < that.grades.length; i++) {
                    if (that.grades[i]["grade"] === that.grade) {
                        that.subjects = that.grades[i]["subject"];
                        break;
                    }
                }
            },
            formatDate(now) {
                var year = now.getFullYear();
                var month = now.getMonth() + 1;
                var date = now.getDate();
                var hour = now.getHours();
                var minute = now.getMinutes();
                var second = now.getSeconds();
                return year + "-" + month + "-" + date + " " + hour + ":" + minute + ":" + second;
            }
        }
    }
</script>

<style scoped>
    .plane {
        width: 70%;
        margin: auto;
    }

    .select {

    }

    .table {
        margin-top: 100px;
    }
</style>
